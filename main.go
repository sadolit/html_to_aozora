package main

import (
"os" 
"fmt"
"flag"
"io/ioutil"
"strings"
)
var inputFilename string
var outputFilename string

func mainExitCode() int {
	flag.StringVar(&inputFilename, "f", "", "Input filename")
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
	fmt.Println(strings.TrimLeft(RemoveHtmlTags(RubyToBunko(AddNewline(string(data)))), "\r\n \t"))
	return 0
}

func main() {
	os.Exit(mainExitCode())
}
