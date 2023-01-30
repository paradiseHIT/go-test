package main

func main() {
	// 创建一个整型通道
  ch := make(chan int)

  // 尝试将0通过通道发送
  ch <- 0
  <- ch
}
