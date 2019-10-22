package main

import (
	"container/heap"
	"fmt"
	"math"
)

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].score < pq[j].score
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*puzzle)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *puzzle, value puzzle, priority int) {
	item.mapa = value.mapa
	item.score = priority
	heap.Fix(pq, item.index)
}

type PriorityQueue []*puzzle

type puzzMap [][]int

type puzzle struct {
	mapa  puzzMap
	side  int
	depth int
	score int
	dad   *puzzle
	index int
}

type coor struct {
	x int
	y int
}

type puzzleSolver struct {
	algh    alghoritmoD
	dict    map[string]int
	heap    PriorityQueue
	maxHeap int
	start   *puzzle
	end     map[int]*coor
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

type alghoritmoD func(start *puzzle, end map[int]*coor) int

func manhatanDistance(start *puzzle, end map[int]*coor) int {
	var sum float64
	for y, h := range start.mapa {
		for x, v := range h {
			st := end[v]
			//fmt.Printf("%d %d \n", x, y)
			sum = math.Abs(float64(st.x-x)) + math.Abs(float64(st.y-y)) + sum
			//fmt.Printf("%f \n", h)
		}
	}
	//fmt.Printf("%f\n", sum)
	return int(sum)
}
func titlesOutOfPlace(start *puzzle, end map[int]*coor) int {
	return 2
}
func euclideanDistance(start *puzzle, end map[int]*coor) int {
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
func checkDict(pu *puzzle, ps *puzzleSolver) bool {
	a := fmt.Sprint(pu.mapa)
	if ps.dict[a] == 0 {
		ps.dict[a] = 1
		return true
	}
	return false
}

func coordPuzzle(end *puzzle) map[int]*coor {
	m := make(map[int]*coor)
	b := end.mapa
	for y, h := range b {
		for x, v := range h {
			m[v] = &coor{
				x: x,
				y: y,
			}
		}
	}
	return m
}
func createSolver(start *puzzle, end *puzzle, algh string) puzzleSolver {

	so := puzzleSolver{
		algh:    addAlgoritm(algh),
		dict:    make(map[string]int),
		heap:    make(PriorityQueue, 1),
		maxHeap: 0,
		start:   start,
		end:     coordPuzzle(end),
	}

	so.heap[0] = start
	return so
}

func (p puzzleSolver) Solve() puzzle {
	heap.Init(&(p.heap))
	t := 0
	for p.heap.Len() > 0 {
		t++
		item := heap.Pop(&p.heap).(*puzzle)
		arrayPuzzles := createArrayPuzzle(item)
		for i, _ := range arrayPuzzles {
			puzzle := arrayPuzzles[i]
			if checkDict(&puzzle, &p) {
				alghD := p.algh(&puzzle, p.end)
				if alghD == 0 {
					return puzzle
				}
				//fmt.Printf("%d \n", alghD)
				puzzle.score = puzzle.depth + alghD
				heap.Push(&p.heap, &puzzle)
				fmt.Printf("%+v \n", puzzle)
			}
		}
		if t == 100 {
			break
		}
	}
	return *p.start
}

func main() {
	a := puzzle{
		mapa:  puzzMap{{0, 2, 3}, {1, 4, 5}, {8, 7, 6}},
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
