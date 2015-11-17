package main

import (
	"flag"
	"fmt"
	"github.com/marcvanzee/satsolver-go/satinstance"
	"github.com/marcvanzee/satsolver-go/solvers"
	"github.com/marcvanzee/satsolver-go/watchlist"
	"os"
)

var verbose = flag.Bool("verbose", true, "Verbose output")
var allSolutions = flag.Bool("all", true, "Output all possible solutions")
var brief = flag.Bool("brief", false, "Only output variables assigned true")
var startingWith = flag.String("starting_with", "", "Only output variables with names starting with the given string")
var iterative = flag.Bool("recursive", false, "Use recursive algorithm instead of iterative")
var file = flag.String("i", "", "Read from given file instead of stdin")

func exit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// no error
}

func solve(s satinstance.SATInstance, it bool, verbose bool) [][]int {
	n := len(s.Vars)
	watchlist := watchlist.NewWatchlist(s)

	if len(watchlist) == 0 {
		return nil
	}

	assignment := make([]int, n, n)

	for i := range assignment {
		assignment[i] = satinstance.NONE
	}

	ret := solvers.NewSolver(it).Solve(s, watchlist, assignment, 0, verbose)

	return ret
}

func main() {
	flag.Parse()

	var err error

	instance := satinstance.NewSATInstance()

	defer func() { exit(err) }()

	if err := instance.Init(file); err != nil {
		return
	}

	assignments := solve(instance, *iterative, *verbose)
	count := 0

	for _, assignment := range assignments {
		if *verbose {
			fmt.Printf("Found satisfying assignment #%v:\n", count)
			fmt.Println(instance.AssignmentToString(assignment, *brief, *startingWith))
		}
		count += 1
		if !*allSolutions {
			break
		}
	}

	if *verbose && count == 0 {
		fmt.Println("No satisfying assignment exists.")
	}
}
