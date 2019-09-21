package main

import (
	"../crawler/engine"
	"../crawler/zhenai/parser"
	"../crawler/scheduler"
	"./persist/client"
	"fmt"
	"./config"
)
func main()  {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d",config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEnine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

//城市列表解释器
//城市解释器
//用户解释器
