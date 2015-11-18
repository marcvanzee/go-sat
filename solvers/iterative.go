package solvers

import (
	"fmt"
	"github.com/marcvanzee/go-solver/satproblem"
)

type IterativeSolver struct{}

func (i IterativeSolver) Solve(s satproblem.SATInstance, w satproblem.Watchlist, ass []int, d int, verbose bool) [][]int {
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
					ass[d] = satproblem.NONE
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
				ass[d] = satproblem.NONE
				d -= 1
			}
		}
	}

	return ret
}
