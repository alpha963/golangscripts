// mytes project main.go
package main

import (
	"fmt"
)

func slicp(x, y int) (int, int) {
	var c int
	c = y - x + 1
	return x - 50000, c / 50000
}

func main() {
	var a, b int = 10110003720001, 10110003920000
	a, c := slicp(a, b)

	p := []int{}
	for i := 0; i < c; i++ {
		a += 50000
		p = append(p, a)
	}

	for i := 0; i < len(p); i++ {
		if p[i] < b {
			fmt.Printf("p[%d] == %d,%d\n",
				i, p[i], p[i]+49999)
		}
	}
}
