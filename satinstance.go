package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NOT = "~"

type SATInstance struct {
	Vars      []string
	Clauses   [][]int
	VarTable  map[string]int
	Watchlist Watchlist
}

func NewSATInstance() SATInstance {
	return SATInstance{
		nil,
		nil,
		make(map[string]int),
		make(map[int]([][]int)),
	}
}

func (s SATInstance) String() string {
	return fmt.Sprintf("variables: %v\nClauses: %v\nVariable_table: %v\nWatchlist: %v",
		s.Vars, s.Clauses, s.VarTable, s.Watchlist)
}

func (s *SATInstance) ParseLine(line string) {
	clause := make([]int, 0, 10)

	for _, literal := range strings.Split(line, " ") {
		negated := 0
		if strings.Index(literal, NOT) == 0 {
			negated = 1
		}

		variable := literal[negated:]

		if _, ok := s.VarTable[variable]; !ok {
			s.VarTable[variable] = len(s.Vars)
			s.Vars = append(s.Vars, variable)
		}

		// to go from the variable to our encoded literal,
		// we multiply by 2 and add 1 if negated.
		encoded_literal := s.VarTable[variable]<<1 | negated
		clause = append(clause, encoded_literal)
	}
	s.Clauses = append(s.Clauses, clause)
}

func (s *SATInstance) ParseFile(f *os.File) {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		} else if strings.Index(line, "#") != 0 {
			s.ParseLine(line)
		}
	}
}

func (s *SATInstance) Init(f *string) (err error) {
	file := os.Stdin
	if *f != "" {
		file, err = os.Open(*f)
		if err != nil {
			return err
		}
	}

	s.ParseFile(file)
	s.Watchlist.Init(s.Clauses)

	return nil
}

func (s *SATInstance) LiteralToString(literal int) string {
	str := ""

	if literal&1 > 0 {
		str = "~"
	}

	return str + s.Vars[literal>>1]
}

func (s *SATInstance) ClauseToString(clause []int) string {

	str := make([]string, len(clause))

	for i, v := range clause {
		str[i] = s.LiteralToString(v)
	}

	return strings.Join(str, " ")
}

func (s *SATInstance) AssignmentToString(assignment []int, brief bool, starting_with string) string {
	literals := make([]string, len(s.Vars))

	for i, v := range s.Vars {
		a := assignment[i]
		if strings.Index(v, starting_with) == 0 {
			if a == 0 && !brief {
				literals = append(literals, NOT+v)
			} else if a == 1 {
				literals = append(literals, v)
			}
		}
	}

	return strings.Join(literals, " ")
}
