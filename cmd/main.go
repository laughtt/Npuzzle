package main

import (
	"container/heap"
	"fmt"
)

type puzzMap [][]int

type puzzle struct {
	mapa  puzzMap
	side  int
	depth int
	score int
	dad   *puzzle
	index int
}

type puzzleSolver struct {
	algh    alghoritmoD
	dict    map[string]int
	heap    PriorityQueue
	maxHeap int
}

func duplicateArray(matrix puzzMap) puzzMap {
	duplicate := make(puzzMap, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]int, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func createPuzzle(p *puzzle, mp puzzMap) puzzle {
	newPuzzle := puzzle{
		mapa:  mp,
		side:  p.side,
		depth: p.depth + 1,
		score: 10,
		dad:   p,
	}
	fmt.Printf("%+v \n", newPuzzle)
	return newPuzzle
}
func createArrayPuzzle(p *puzzle) []puzzle {
	l := make([]puzzle, 0)

	for i := 0; i < p.side; i++ {
		for y := 0; y < p.side; y++ {
			if p.mapa[i][y] == 0 {
				if i > 0 {
					top := duplicateArray(p.mapa)
					value := top[i-1][y]
					top[i-1][y] = 0
					top[i][y] = value
					newPuzzle := createPuzzle(p, top)
					l = append(l, newPuzzle)
				}
				if i < p.side-1 {
					down := duplicateArray(p.mapa)
					value := down[i+1][y]
					down[i+1][y] = 0
					down[i][y] = value
					newPuzzle := createPuzzle(p, down)
					l = append(l, newPuzzle)
				}
				if y > 0 {
					left := duplicateArray(p.mapa)
					value := left[i][y-1]
					left[i][y-1] = 0
					left[i][y] = value
					newPuzzle := createPuzzle(p, left)
					l = append(l, newPuzzle)
				}
				if y < p.side-1 {
					right := duplicateArray(p.mapa)
					value := right[i][y+1]
					right[i][y+1] = 0
					right[i][y] = value
					newPuzzle := createPuzzle(p, right)
					l = append(l, newPuzzle)
				}
				return l
			}
		}
	}
	return l
}

type alghoritmoD func(start *puzzle, end *puzzle) int

func manhatanDistance(start *puzzle, end *puzzle) int {
	return 1
}
func titlesOutOfPlace(start *puzzle, end *puzzle) int {
	return 2
}
func euclideanDistance(start *puzzle, end *puzzle) int {
	return 3
}

func addAlgoritm(s string) alghoritmoD {

	switch s {
	case "mh":
		a := manhatanDistance
		return a
	case "to":
		a := titlesOutOfPlace
		return a
	case "ed":
		a := euclideanDistance
		return a
	default:
		a := manhatanDistance
		return a
	}
}
func checkDict(p *puzzle) bool {
	return true
}
func createSolver(start *puzzle, end *puzzle, algh string) puzzleSolver {
	so := puzzleSolver{
		algh:    addAlgoritm(algh),
		dict:    make(map[string]int),
		heap:    make(PriorityQueue, 1),
		maxHeap: 0,
	}
	so.heap[0] = start
	return so
}

func (p puzzleSolver) Solve() int {
	heap.Init(&(p.heap))
	return 1
}

func main() {
	a := puzzle{
		mapa:  puzzMap{{0, 2, 3}, {4, 5, 6}, {7, 8, 1}},
		side:  3,
		depth: 0,
		score: 10,
		dad:   nil,
		index: 0,
	}
	b := puzzle{
		mapa:  puzzMap{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}},
		side:  3,
		depth: 0,
		score: 0,
		dad:   nil,
		index: 0,
	}
	mh := "mh"
	solver := createSolver(&a, &b, mh)
	fmt.Println(solver.Solve())
	//arrayPuzzles := createArrayPuzzle(&a)

	//fmt.Printf("%v %+v", a, a)
}
