package shared

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"
)

func Open(file string, f func(fileScanner *bufio.Scanner)) {
	readFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				fmt.Println(string(debug.Stack()))
				return
			}
		}()
		f(fileScanner)
	}
	readFile.Close()
}

func Run(f func() any) {
	result := f()
	fmt.Println(fmt.Sprint("The result is : ", result))
}
