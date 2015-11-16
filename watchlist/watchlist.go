package watchlist

import (
	"fmt"
	"github.com/marcvanzee/satsolver-go/satinstance"
	"strings"
)

type Watchlist map[int]([][]int)

func NewWatchlist(s SATInstance) Watchlist {
	var w Watchlist
	w = make(map[int]([][]int))

	for _, c := range s.Clauses {
		w[c[0]] = append(w[c[0]], c)
	}

	return w
}

func (w Watchlist) String(s SATInstance) string {
	ret := ""
	for l, w2 := range w {
		lStr := s.LiteralToString(l)
		var cStr []string
		for _, c := range w2 {
			cStr = append(cStr, s.ClauseToString(c))
		}
		ret += lStr + ": [" + strings.Join(cStr, ", ") + "]\n"
	}

	return ret
}

// update the watchlist when literal fl has just been set to false
// do this by make any clause that is currently watching fl watch another literal
// returns false if this is impossible, which means that all literals in a clause have
// already been assigned flase.
func (w Watchlist) Update(s SATInstance, fl int, assignment []int, verbose bool) bool {

	// continue as long as there are still clauses watching over fl
	for {
		if _, ok := w[fl]; ok {
			found := false

			// select the first clause that is currently watching fl
			clause := w[fl][0]

			// try to watch another literal of this clause
			for _, alt := range clause {
				v := alt >> 1 // v is the index of the literal
				a := alt & 1  // a == 1: alt is uneven, i.e. alt is false
				// a == 0: alt is even, i.e. alt is true

				// if alt is unassigned or true, we can watch it
				if assignment[v] == NONE || assignment[v] == a^1 {
					found = true

					if len(w[fl]) > 1 {
						w[fl] = w[fl][1:]
					} else {
						delete(w, fl)
					}
					w[alt] = append(w[alt], clause)
					break
				}
			}

			if !found {
				if verbose {
					fmt.Println("Current watchlist:",
						w.String(s))
					fmt.Println("Current assignment:",
						s.AssignmentToString(assignment, false, ""))
					fmt.Println("Clause",
						s.ClauseToString(clause),
						"contradicted")
				}
				return false
			}
		} else {
			// all clauses that were watching fl have been reassigned
			return true
		}
	}
}
