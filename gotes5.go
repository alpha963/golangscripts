package main
//switch 星期几 操作系统信息；当前时间小时数。
import (
	"fmt"
	"runtime"
	"time"

)

func main() {

	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("\n When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("The day after tomorrow.")
	case today + 3:
		fmt.Println("Three days from now.")
	case today + 4:
		fmt.Println("Four days from now.")
	case today + 5:
		fmt.Println("Five days from now.")
	default:
		fmt.Println("Six days from now.")
	}
	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	
        baseTime := time.Date(1980, 1, 6, 0, 0, 0, 0, time.UTC)
        date := baseTime.Add(4*time.Month)
        fmt.Println(date)


	
 

}