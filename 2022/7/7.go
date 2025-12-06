package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

var rootFolder = &folder{
	fullPath:     "",
	files:        []file{},
	folders:      map[string]*folder{},
	parentFolder: nil,
}

var currentFolder = rootFolder

var fullPathToDirectory = map[string]*folder{
	"/": rootFolder,
}

type file struct {
	name string
	size int64
}

type folder struct {
	fullPath     string
	files        []file
	folders      map[string]*folder
	parentFolder *folder
}

func (f *folder) size() int64 {
	result := int64(0)
	for _, file := range f.files {
		result += file.size
	}
	for _, folder := range f.folders {
		result += folder.size()
	}
	return result
}

func (f *folder) cd(subFolderLine string) {
	if subFolderLine == "$ cd .." {
		currentFolder = currentFolder.parentFolder
		return
	}
	if subfolder, ok := fullPathToDirectory[fmt.Sprint(f.fullPath, "/", subFolderLine[5:])]; ok {
		currentFolder = subfolder
		return
	}
	panic(fmt.Sprint("a case where the cd is passed on unknown folder occurs: ", subFolderLine))
}

func (f *folder) listDirectory(directoryLine string) {
	subFolderPath := fmt.Sprint(f.fullPath, "/", directoryLine[4:])
	newFolder := &folder{
		fullPath:     subFolderPath,
		files:        []file{},
		folders:      map[string]*folder{},
		parentFolder: f,
	}
	fullPathToDirectory[subFolderPath] = newFolder
	currentFolder.folders[subFolderPath] = newFolder
}

func (f *folder) listFile(fileLine string) {
	fileElements := strings.Split(fileLine, " ")
	size, _ := strconv.ParseInt(fileElements[0], 10, 64)
	currentFolder.files = append(currentFolder.files, file{
		name: fileElements[1],
		size: size,
	})
}

func main() {
	shared.Run(func() any {

		shared.Open("7.txt", func(fileScanner *bufio.Scanner) {
			line := fileScanner.Text()
			if line == "$ ls" || line == "$ cd /" {
			} else if line[:4] == "$ cd" {
				currentFolder.cd(line)
			} else if line[:4] == "dir " {
				currentFolder.listDirectory(line)
			} else {
				currentFolder.listFile(line)
			}

		})
		count := int64(0)
		for _, folder := range fullPathToDirectory {
			currentSize := folder.size()
			if currentSize <= 100000 {
				count += currentSize
			}
		}
		return count
	})
}
