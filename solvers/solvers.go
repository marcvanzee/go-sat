package solvers

import (
	"fmt"
	"github.com/marcvanzee/satsolver-go/satinstance"
	"github.com/marcvanzee/satsolver-go/watchlist"
)

func SolveRecursive(s SATInstance, w Watchlist, ass []int, d int, verbose bool) [][]int {
	if d == len(s.Vars) {
		return [][]int{ass}
	}

	var ret [][]int

	for _, a := range []int{0, 1} {
		if verbose {
			fmt.Printf("Trying %d = %d\n", s.Vars[d], a)
		}
		ass[d] = a
		if w.Update(s, (d<<1)|a, ass, verbose) {
			sol := solve(s, w, ass, d+1, verbose)

			ret = append(ret, sol...)
		}
	}
	ass[d] = satinstance.None

	return ret
}

func SolveIterative() {
}
