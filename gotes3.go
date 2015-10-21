package main
//勾股定理。

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	var x, y,sum,i int = 3, 4,0,1
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)
	fmt.Printf("f and z is of type %T,%T\n", f,z)

	for ;i<=100;i++ {
	sum+=i
	}
	fmt.Println(sum)
	fmt.Println(sqrt(2), sqrt(-4))
}