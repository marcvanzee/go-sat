package main

/**
* ===== GO-SAT - SIMPLE SAT SOLVER IN GO =====
* == By Marc van Zee (marcvanzee@gmail.com) ==
* ============================================
* =     www.github.com/marcvanzee/go-sat/    =
* =                                          =
* ============================================
*
* View a more detailed explanation of this solver in Python:
* http://sahandsaba.com/understanding-sat-by-implementing-a-simple-sat-solver-in-python.html
*
*
* Run the SAT solver from here with optional arguments:
* -all
*       Output all possible solutions (default true)
*  -brief
*        Only output variables assigned true
*  -i string
*        Read from given file instead of stdin
*  -recursive
*        Use recursive algorithm instead of iterative
*  -starting_with string
*        Only output variables with names starting with the given string
*  -verbose
*        Verbose output (default true)
 */

import (
	"flag"
	"fmt"
	"github.com/marcvanzee/go-solver/satproblem"
	"github.com/marcvanzee/go-solver/solvers"
	"os"
)

var verbose = flag.Bool("verbose", true, "Verbose output")
var allSolutions = flag.Bool("all", true, "Output all possible solutions")
var brief = flag.Bool("brief", false, "Only output variables assigned true")
var startingWith = flag.String("starting_with", "", "Only output variables with names starting with the given string")
var recursive = flag.Bool("recursive", false, "Use recursive algorithm instead of iterative")
var file = flag.String("i", "", "Read from given file instead of stdin")

func exit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// no error
}

func solve(s satproblem.SATInstance) [][]int {

	n := len(s.Vars)

	// initialize the watchlist, where each clauses in the SATInstance initially
	// watches it first literal
	watchlist := satproblem.NewWatchlist(s)

	if len(watchlist) == 0 {
		return nil
	}

	// initially we do not assignment a truth value to any of the literals
	assignment := make([]int, n, n)

	for i := range assignment {
		assignment[i] = satproblem.NONE
	}

	ret := solvers.NewSolver(*recursive).
		Solve(s, watchlist, assignment, 0, *verbose)

	return ret
}

func main() {
	flag.Parse()

	var err error
	defer func() { exit(err) }()

	instance := satproblem.NewSATInstance()

	// initalize tries to read from file, if file is empty it reads from stdin
	if err := instance.Init(file, *verbose); err != nil {
		return
	}

	// solve returns a slide of solution assignments,
	// which assign truth values to literals
	assignments := solve(instance)

	// print the assignments
	count := 0
	for _, assignment := range assignments {
		if *verbose {
			fmt.Printf("Found satisfying assignment #%v: ", count)
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
