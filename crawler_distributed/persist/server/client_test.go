package main

import (
	"../../../crawler/engine"
	"../../../crawler/model"
	"../../persist"
	"../../rpcsupport"
	"go-practice/crawler_distributed/config"
	"gopkg.in/olivere/elastic.v5"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":12345"
	//start ItemSaverServer
	go serveRpc1(host, "test1")

	time.Sleep(time.Second)
	//start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1972182204",
		Type: "zhenai",
		Id:   "1972182204",
		Payload: model.Profile{
			Name:      "磕磕",
			Age:       43,
			Height:    161,
			Income:    "5001-8000元",
			Marriage:  "离异",
			Education: "大学本科",
			Hokou:     "阿坝",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
	//Call save
}

func serveRpc1(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.199.102:9200"),
		elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
