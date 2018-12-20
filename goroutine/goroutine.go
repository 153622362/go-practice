package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) { //协程  轻量级"线程"  非抢占式 协程主动交出控制权 编译器/解释器/虚拟机层面的多任务 多个协程在一个线程或多个线程上进行 子程序是协程的一个特例
		//goroutine可能切换点 I/O select channel 等待锁 函数调用(有时) runtime.Gosched()
			for {
				fmt.Printf("Hello from "+
					"goroutine %d\n", i)
				//runtime.Gosched() 手动交出控制权
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
