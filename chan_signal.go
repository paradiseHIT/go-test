package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var stopCh chan struct{}

func run() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	stopCh = make(chan struct{})
	go sig(c)
	loop()
	fmt.Println("done")
}

func sig(c chan os.Signal) {
	fmt.Println("sigin")
	<-c
	fmt.Println("sigout")
	stopCh <- struct{}{}
}

func loop() {
	for i := 1; i <= 100; i++ {
		select {
		case <-stopCh:
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	run()
}
