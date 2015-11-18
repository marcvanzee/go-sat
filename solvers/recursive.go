package solvers

import (
	"fmt"
	"github.com/marcvanzee/go-solver/satproblem"
)

type RecursiveSolver struct{}

func (r RecursiveSolver) Solve(s satproblem.SATInstance, w satproblem.Watchlist, ass []int, d int, verbose bool) [][]int {

	if d == len(s.Vars) {
		ret := make([]int, len(ass))
		copy(ret, ass)
		return [][]int{ret}
	}

	var ret [][]int

	for a := range []int{0, 1} {
		if verbose {
			fmt.Printf("Trying %v = %v\n", s.Vars[d], a)
		}
		ass[d] = a
		if w.Update(s, (d<<1)|a, ass, verbose) {
			ret = append(ret, r.Solve(s, w, ass, d+1, verbose)...)
		}
	}
	ass[d] = satproblem.NONE

	return ret
}
