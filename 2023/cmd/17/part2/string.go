package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (g grid) stringPath(goal *cell, parent map[*cell]*cell) string {
	stringGrid := make([][]string, len(g))
	for i, row := range g {
		stringGrid[i] = make([]string, len(row))
		for j, n := range row {
			stringGrid[i][j] = strconv.Itoa(n[0].loss)
		}
	}

	for current, ok := goal, true; ok; current, ok = parent[current] {
		stringGrid[current.index.i][current.index.j] = "*"
	}

	stringRows := make([]string, len(stringGrid))
	for i, row := range stringGrid {
		stringRows[i] = strings.Join(row, "")
	}

	s := strings.Join(stringRows, fmt.Sprintln())

	s += fmt.Sprintln()
	s += fmt.Sprintln()

	for current, ok := goal, true; ok; current, ok = parent[current] {
		s += fmt.Sprintf("%v\n", current)
	}

	return s
}
