package main

type Watchlist map[int]([][]int)

func (w Watchlist) Init(clauses [][]int) {
	for _, c := range clauses {
		w[c[0]] = append(w[c[0]], c)
	}
}

func (w Watchlist) Update(s SATInstance, fl int, assignment []int, verbose bool) {
}
