package solvers

import (
	"github.com/marcvanzee/go-solver/satproblem"
)

func NewSolver(recursive bool) Solver {
	if !recursive {
		return IterativeSolver{}
	} else {
		return RecursiveSolver{}
	}
}

type Solver interface {
	Solve(satproblem.SATInstance, satproblem.Watchlist, []int, int, bool) [][]int
}
