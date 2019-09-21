package main

import (
	"context"
	"fmt"
	"time"
	"unsafe"
)

func paramc(param ...int) {
	//for _,c := range param {
	//	fmt.Println(c)
	//}
}

type Node struct {
	Value       int
	Left, Right *Node
}

//结构定义方法
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func generator() <-chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			out <- i
			i++
		}
	}()
	return out
}

func main() {

	var c1 = generator()

	for v := range c1 { //遍历channel
		fmt.Println(v)
		if v == 8 {
			break
		}
	}
	fmt.Println("good job")
	//str,_ := ioutil.ReadFile("errors.log")
	//fmt.Printf("%s", str)
	bytes := []byte("I am byte array !")
	str := (*string)(unsafe.Pointer(&bytes))
	bytes[0] = 'i' //修改字符串
	fmt.Println(*str)

	//var root Node
	//
	//root.Left = CreateNode(2)
	//root.Right = CreateNode(4)
	//
	//root.Left.Print()
	//root.Right.Print()
	//root.SetValue(3)
	//root.Print()

	//Make a background context
	//ctx := context.Background()
	//Derive a context with cancel
	//ctxWithCancel, cancelFunction := context.WithCancel(ctx)
	//go doContext(ctxWithCancel, cancelFunction)
	//cancelFunction()
	//time.Sleep(time.Second * 5)

	//paramc(1,2,3,4)

	//c := make(chan int)  // 分配一个信道
	//// 在Go程中启动排序。当它完成后，在信道上发送信号。
	//go func() {
	//	c <- 1  // 发送信号，什么值无所谓。
	//}()
	//<-c   // 等待排序结束，丢弃发来的值。

	//Server()
}

type Request struct {
	args       int
	resultChan chan int
}

func testp(i []int) {
	fmt.Println(i)
}

var req [10]Request

func Server() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		req[i] = Request{
			args:       1,
			resultChan: c,
		}
	}

	go handle()
	for {
		select {
		case n := <-c:
			fmt.Println(n)
		}
	}

	time.Sleep(time.Second * 5)

}

var i int

func handle() {
	for {
		fmt.Println("test")
		chanl := req[0]
		chanl.resultChan <- 333
		i++

	}

}

func doContext(ctx context.Context, cancelFunc context.CancelFunc) {
	for {
		fmt.Println(1)
		//cancelFunc()
		//go func() {
		//	println(22222)
		//}()

		select {
		case <-ctx.Done():
			fmt.Println("66666")
			return
		}
		time.Sleep(time.Second)
	}

}
