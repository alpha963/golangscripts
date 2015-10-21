package main
//勾股定理。

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)
	fmt.Printf("f and z is of type %T,%T\n", f,z)

	sum:=0
	for i:=0;i<=101;i++{
	sum+=i
	}
	fmt.Println(sum)
}