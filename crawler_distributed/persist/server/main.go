package main

import (
	"../../rpcsupport"
	"../../persist"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"../../config"
)
func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d",config.ItemSaverPort), config.ElasticIndex))
}

//index elasticsearch数据库索引
//RPC服务
func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.199.102:9200"),
		elastic.SetSniff(false))
	if err != nil {
		return  err
	}
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client:client,
			Index:index,
		})
}
