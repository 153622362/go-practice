package persist

import (
	"../model"
	"../persist"
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"testing"
	"../engine"
)


func TestSave(t *testing.T) {

	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetURL("http://192.168.199.102:9200"),
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}
	expected := engine.Item{
		Url : "http://album.zhenai.com/u/1972182204",
		Type: "zhenai",
		Id: "1972182204",
		Payload:model.Profile{
		Name:		"磕磕",
		Age:		43,
		Height:		161,
		Income:		"5001-8000元",
		Marriage:	"离异",
		Education:	"大学本科",
		Hokou:		"阿坝",
		},
	}

	// Save expected item
	err = persist.Save(client,"dating_profile", expected)

	if err != nil {
		panic(err)
	}

	////  TODO: Try to start up elastic search
	//// here using docker go client.
	//client, err := elastic.NewClient(elastic.SetURL("http://192.168.199.102:9200"),
	//	elastic.SetSniff(false))
	//
	//if err != nil {
	//	panic(err)
	//}

	// Fetch saved item
	resp, err := client.Get().
		Index("dating_profile").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	//t.Logf("%+v", resp)
	t.Logf("%s", *resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)

	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; profile %v", actual, expected)
	}
}