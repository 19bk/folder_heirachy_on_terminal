package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func printDirStructure(path string, prefix string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		fmt.Printf("%s|-- %s\n", prefix, file.Name())
		if file.IsDir() {
			printDirStructure(filepath.Join(path, file.Name()), prefix+"|   ")
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory>")
		return
	}
	dir := os.Args[1]
	fmt.Println(dir + "/")
	printDirStructure(dir, "")
}
