package solvers

import (
	"fmt"
	"github.com/marcvanzee/satsolver-go/satinstance"
	"github.com/marcvanzee/satsolver-go/watchlist"
)

func NewSolver(it bool) Solver {
	if it {
		return IterativeSolver{}
	} else {
		return RecursiveSolver{}
	}
}

type Solver interface {
	Solve(satinstance.SATInstance, watchlist.Watchlist, []int, int, bool) [][]int
}

type RecursiveSolver struct{}
type IterativeSolver struct{}

func (r RecursiveSolver) Solve(s satinstance.SATInstance, w watchlist.Watchlist, ass []int, d int, verbose bool) [][]int {

	if d == len(s.Vars) {
		ret := make([][]int, 1)
		ret[0] = make([]int, len(ass))
		copy(ret[0], ass)
		return ret
	}

	var ret [][]int

	for _, a := range []int{0, 1} {
		if verbose {
			fmt.Printf("Trying %v = %v\n", s.Vars[d], a)
		}
		ass[d] = a
		if w.Update(s, (d<<1)|a, ass, verbose) {
			ret = append(ret, r.Solve(s, w, ass, d+1, verbose)...)
		}
	}
	ass[d] = satinstance.NONE

	return ret
}

func (i IterativeSolver) Solve(s satinstance.SATInstance, w watchlist.Watchlist, ass []int, d int, verbose bool) [][]int {
	return nil
}
