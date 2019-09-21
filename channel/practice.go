package main

import (
	"fmt"
	"sync"
	"time"
)

type worker1 struct {
	in   chan int
	done func()
}

func doWorker1(id int,
	w worker1) {
	for n := range w.in {
		//n, ok := <-c
		//if !ok {
		//	break
		//}
		fmt.Printf("Worker %d received %c \n", id, n)
		w.done()
	}
}

//chan<- 告诉别人这是用来发数据的
//<-chan 告诉别人只能用来收数据
//chan 告诉别人收发均可
func createWorker1(id int, wg *sync.WaitGroup) worker1 {
	w := worker1{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}

	go doWorker1(id, w)

	return w

}

//函数和channel都能作为参数和返回值
func channelDemo() {
	var wg sync.WaitGroup
	var workers [10]worker1 //只能发数据
	for i := 0; i < 10; i++ {
		workers[i] = createWorker1(i, &wg)
	}

	wg.Add(10)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	wg.Wait() //等待Goroutine执行完成
}

//func bufferChannel()  {
//	c := make(chan int, 3)
//	go worker1(0, c)
//	c <- 'a'
//	c <- 'b'
//	c <- 'c'
//	close(c) //可关闭
//	time.Sleep(time.Millisecond)
//}

//理论基础CSP模型 communication Sequential Process
//不要通过共享内存来通信  通过通信来共享内存
//如何知道把事情做完了呢  共享一个flag 设置 true你就知道做完了
//如何知道Println完毕了呢?
func main() {
	//channelDemo()
	//test1()
	//bufferChannel()
	defer2()

	defer1()
}

func defer1() {
	defer func() {
		fmt.Println(2)
	}()
	fmt.Println(1)

}

func defer2() {
	defer func() {
		fmt.Println(4)
	}()
	fmt.Println(3)

}

func test1() {
	c1 := make(chan int)
	//var c1  chan int
	go func() {
		for {
			c1 <- 1
		}
	}()
	for {
		select {
		case n := <-c1:
			fmt.Println(n)
		}
	}

	time.Sleep(time.Millisecond)
}

func select1() {
	var c1, c2 chan int // c1 and c2 = nil

	select {
	case n := <-c1:
		fmt.Println("Received from c1:", n)
	case n := <-c2:
		fmt.Println("Received from c2:", n)
	default:
		fmt.Println("No Value received")
	}
}
