package main
//使用时间做种子生产随机数。
import ("fmt"
)

func add(x ,y int) int {
return x*y
}

func main() {

a,b,i:=5,4,3
python,java:=true,"no!"

var c int =a-b+1
a,b=add(3,2),add(4,5)

fmt.Println("距离地面高度为：",add(a,i)-add(c,i),"（米）")
fmt.Println(python,java,a,b)

}