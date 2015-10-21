package main
//使用时间做种子生产随机数。
import ("fmt"
"math/rand"
"time"
)

func main() {
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<10;i++{
	fmt.Println(r.Intn(100))
	}

}