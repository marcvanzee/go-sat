package main

import (
	"flag"
	"fmt"
	"os"
)

var verbose = flag.Bool("verbose", true, "Verbose output")
var allSolutions = flag.Bool("all", true, "Output all possible solutions")
var brief = flag.Bool("brief", true, "Only output variables assigned true")
var startingWith = flag.String("starting_with", "", "Only output variables with names starting with the given string")
var iterative = flag.Bool("iterative", true, "Use the iterative algorithm")
var file = flag.String("i", "", "Read from given file instead of stdin")

func exit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// no error
}

func main() {
	flag.Parse()

	var err error

	instance := NewSATInstance()

	defer func() { exit(err) }()

	if err := instance.Init(file); err != nil {
		return
	}

	fmt.Println(instance)
}
