package main

import "os"
import "fmt"
import "flag"

var inputFilename string
var outputFilename string

func mainExitCode() int {
	flag.StringVar(&inputFilename, "f", "", "Input filename")
	flag.StringVar(&outputFilename, "o", "", "Output filename")
	flag.Parse()
	if inputFilename == "" {
		fmt.Println("Input Filename can't be empty, print --help for usage")
		return 1
	}
	fmt.Println("Input file: [" + inputFilename + "]")
	fmt.Println("Output file: [" + outputFilename + "]")
	return 0
}

func main() {
	os.Exit(mainExitCode())
}
