package main

import "fmt"

type Vertex struct {
	X int
	Y int
}


func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println("i=:",i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println(Vertex{1, 2})
	v:=Vertex{5,6}
	v.X=9
	fmt.Println(v.X,v.Y)
	z := &v
	z.Y= 1e9
	fmt.Println(v)
}