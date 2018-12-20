package persist

import (
	"context"
	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"../engine"
)

func ItemSaver() (chan engine.Item, error) {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetURL("http://192.168.199.102:9200"),
		elastic.SetSniff(false))

	if err != nil {
		return  nil, err
	}

	out := make(chan engine.Item)
	go func() {
		 	itemCount := 0
			for {
				item := <- out
				log.Printf("Item Saver: got item" +
					"#%d: %v", itemCount, item)
				itemCount++

				err := Save(client,"dating_profile", item)
				if err != nil {
					log.Print("Item Sver: error " +
						"saving item %v: %v", item, err)
				}
			}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string,item engine.Item)  error {
	if item.Type == "" {
		return  errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.
		Do(context.Background())

	if err != nil {
		return  err
	}
	//fmt.Printf("%+v", resp)
	return  nil
}