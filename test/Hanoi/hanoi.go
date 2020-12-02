package main

import (
	"flag"
	"fmt"
)

// go run hanoi.go -level=3
func main() {
	level := flag.Int("level", 3, "level=3, level > 0")
	flag.Parse()

	fmt.Println("Hanoi by recursive(", *level, "):")
	Hanoi_Recursive(*level, "A", "B", "C")

	// fmt.Println("")
	// fmt.Println("Hanoi by non-recursive:")
	// Hanoi_Nonrecursive(*level, "A", "B", "C")
}

// Recursive
func Hanoi_Recursive(level int, src, aux, dist string) {
	if level < 1 {
		return
	} else if level == 1 {
		fmt.Println("Move ", src, " To ", dist)
		return
	}

	Hanoi_Recursive(level-1, src, dist, aux)
	Hanoi_Recursive(1, src, aux, dist)
	Hanoi_Recursive(level-1, aux, src, dist)
}

// Non-Recursive
type StackItem struct {
	level int
	src   string
	aux   string
	dist  string
	pnext *StackItem
}

// Non-recursive
func Hanoi_Nonrecursive(level int, src, aux, dist string) {
	if level < 1 {
		fmt.Println("Hanoi, input level is ERROR!")
		return
	}

	// init
	top := &StackItem{} // stack top

	// Stack Push
	top.pnext = &StackItem{level, src, aux, dist, nil}

	for top.pnext != nil {
		// Stack Pop
		t := top.pnext
		top.pnext = t.pnext

		if t.level == 1 {
			fmt.Println("Move ", t.src, " To ", t.dist)
		} else {
			// Stack Push
			top.pnext = &StackItem{t.level - 1, t.aux, t.src, t.dist, top.pnext}
			top.pnext = &StackItem{1, t.src, t.aux, t.dist, top.pnext}
			top.pnext = &StackItem{t.level - 1, t.src, t.dist, t.aux, top.pnext}
		}

	}

}
