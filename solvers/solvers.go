package solvers

import (
	"fmt"
	"github.com/marcvanzee/satsolver-go/satinstance"
	"github.com/marcvanzee/satsolver-go/watchlist"
)

func NewSolver(recursive bool) Solver {
	if !recursive {
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
	ass[d] = satinstance.NONE

	return ret
}

func (i IterativeSolver) Solve(s satinstance.SATInstance, w watchlist.Watchlist, ass []int, d int, verbose bool) [][]int {
	n := len(s.Vars)
	state := make([]int, n)

	var ret [][]int

	for true {
		if d == n {
			assCopy := make([]int, len(ass))
			copy(assCopy, ass)
			ret = append(ret, assCopy)
			d -= 1
			continue
		}

		tried := false
		for a := range []int{0, 1} {
			if (state[d]>>uint(a))&1 == 0 {
				if verbose {
					fmt.Printf("Trying %v = %v\n", s.Vars[d], a)
				}
				tried = true

				state[d] |= 1 << uint(a)
				ass[d] = a
				if !w.Update(s, (d<<1)|a, ass, verbose) {
					ass[d] = satinstance.NONE
				} else {
					d += 1
					break
				}
			}
		}

		if !tried {
			if d == 0 {
				return ret
			} else {
				state[d] = 0
				ass[d] = satinstance.NONE
				d -= 1
			}
		}
	}

	return ret
}
