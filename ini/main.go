package main

import (
	"fmt"
	"github.com/go-ini/ini"
)

type Test struct {
	Aaa string `ini:"aaa"`
}

var file ini.File

func main() {
	file, _ := ini.Load("./ini/conf.ini")

	var test = &Test{}
	file.Section("test").MapTo(test)
	fmt.Println(test.Aaa)
	//var c = make(chan int)
	//var arr [5]int
	//mapTo(&arr)
	//go func( c chan int) {
	//	for {
	//		data := <- c

	//		fmt.Println(data ,runtime.NumGoroutine())
	//	}
	//
	//}(c)
	//
	//for i := 0; i < 100 ; i++  {
	//	c <- i
	//}
	//time.Sleep(time.Second * 5)

}
