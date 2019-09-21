package main

import (
	"fmt"
	"go-practice/rabbitmq/rabbitmq"
)

type TestPro struct {
	msgContent string
}

// 实现发送者
func (t *TestPro) MsgContent() string {
	return t.msgContent
}

// 实现接收者
func (t *TestPro) Consumer(dataByte []byte) error {
	fmt.Println(string(dataByte))
	return nil
}

var mq *rabbitmq.RabbitMQ

//func generator() chan interface{} {
//	out := make(chan interface{})
//
//	go func() {
//		for {
//			c1 := <-out
//			mq.ListenProducer()
//		}
//	}()
//	return out
//}

func main() {
	msg := fmt.Sprintf("这是测试任务")
	t := &TestPro{
		msg,
	}
	queueExchange := &rabbitmq.QueueExchange{
		"test.rabbit",
		"rabbit.key",
		"test.rabbit.mq",
		"direct",
	}
	mq = rabbitmq.New(queueExchange)

	mq.MqConnect()
	//for j := 0; j < 5; j++ {
	//	arr[j] = generator()
	//}

	for i := 0; i < 10000000; i++ {
		mq.RegisterProducer(t)

	}
	mq.Start()

	//i:=0
	//for {
	//	k := i % 5
	//	arr[k] <- "test"
	//	i++
	//}
	//mq.RegisterReceiver(t)

}
