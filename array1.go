package main

import ("fmt"
"math"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	var l,a float64=2,2
	
	f:=math.Pow(l,a)
	fmt.Println(f)

	var i, j, k, count int = 0, 0, 0, 0
    for i = 101; i <= 200; i++ {
        k = int(math.Sqrt(float64(i)))
        for j = 2; j <= k; j++ {
            if i%j == 0 {
                break
            }
        }
        if j == k+1 {
            fmt.Println(i)
            count++
        }
    }
    fmt.Println("total", count)
}
