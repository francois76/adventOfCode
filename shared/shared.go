package shared

import (
	"bufio"
	"os"
)

func Open(file string, f func(fileScanner *bufio.Scanner)) {
	readFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		f(fileScanner)
	}
	readFile.Close()
}
