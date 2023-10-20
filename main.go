package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	
)

func printDirStructure(path string, prefix string, all bool) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if !all && strings.HasPrefix(file.Name(), ".") {
			continue
		}

		fmt.Printf("%s|-- %s\n", prefix, file.Name())
		if file.IsDir() {
			printDirStructure(filepath.Join(path, file.Name()), prefix+"|   ", all)
		}
	}
}

func printHelp() {
	helpText := `
Usage: go run main.go [OPTIONS] <directory>

Options:
  -a, --all    Include hidden files and directories
  -h, --help   Show this help message and exit

Example:
  go run main.go /path/to/directory
  go run main.go --all /path/to/directory
`
	fmt.Println(helpText)
}

func main() {
	allPtr := flag.Bool("all", false, "Include hidden files")
	aPtr := flag.Bool("a", false, "Include hidden files (shorthand)")
	helpPtr := flag.Bool("help", false, "Show help message")
	hPtr := flag.Bool("h", false, "Show help message (shorthand)")

	flag.Parse()

	if *helpPtr || *hPtr {
		printHelp()
		return
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Error: Missing directory argument")
		printHelp()
		return
	}

	dir := args[0]
	all := *allPtr || *aPtr

	fmt.Println(dir + "/")
	printDirStructure(dir, "", all)
}
