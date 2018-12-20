package main

import (
	"./engine"
	"./zhenai/parser"
	"./scheduler"
	"./persist"
)
func main()  {
	// 去重
	// 重构
	itemChan, err := persist.ItemSaver()
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
