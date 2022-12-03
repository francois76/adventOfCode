package shared

import (
	"bufio"
	"fmt"
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

func Run(f func() interface{}) {
	result := f()
	fmt.Println(fmt.Sprint("The result is : ", result))
}
