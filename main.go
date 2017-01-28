package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var inputFilename string
var author string
var title string

func mainExitCode() int {
	flag.StringVar(&inputFilename, "f", "", "Input filename")
	flag.StringVar(&author, "a", "Unknown author", "Author. Can be empty")
	flag.StringVar(&title, "t", "Unknown title", "Title. Can be empty")
	flag.Parse()
	if inputFilename == "" {
		fmt.Println("Input Filename can't be empty, print --help for usage")
		return 1
	}
	data, err := ioutil.ReadFile(inputFilename) // For read access.
	if err != nil {
		fmt.Println(err)
		return 1
	}
	if author != "" && title != "" {
		fmt.Println(title)
		fmt.Println(author)
		fmt.Println("")
	}
	fmt.Println(strings.TrimLeft(RemoveHTMLTags(RubyToBunko(AddNewline(string(data)))), "\r\n \t"))
	return 0
}

func main() {
	os.Exit(mainExitCode())
}
