package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) { //channel可以作为参数 go 函数 channel是一等公民  //接收channel数据
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int //每个都是channel
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ { //为什么打印出来是乱序的Printf调度
		channels[i] <- 'a' + i //分发channel数据
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3) //创建一个channel c  3是缓存 若是没有缓存 必须有channel方接收方
	go worker(0, c)
	c <- 'a' //写入channel
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//fmt.Println("Channel as first-class citizen")
	//chanDemo()
	//fmt.Println("Buffered channel")
	//bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
