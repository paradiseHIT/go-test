package main
import (
    "fmt"
    "time"
)
func main() {
    // 构建一个通道
    ch := make(chan int)
    // 开启一个并发匿名函数
    go func() {
        fmt.Println("start goroutine")
        // 通过通道通知main的goroutine
        time.Sleep(3*time.Second)
				ch <- 0
        fmt.Println("exit goroutine")
    }()
    fmt.Println("wait goroutine")
    // 等待匿名goroutine
		data,ok := <- ch
		fmt.Println(data)
		fmt.Println(ok)

    fmt.Println("all done")
}
