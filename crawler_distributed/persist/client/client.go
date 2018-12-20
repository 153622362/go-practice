package client

import (
	"../../../crawler/engine"
	"../../rpcsupport"
	"log"
	"../../config"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: got item" +
				"#%d: %v", itemCount, item)
			itemCount++


			//Call RPC to save item
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			//err := Save(client,"dating_profile", item)
			if err != nil {
				log.Print("Item Sver: error " +
					"saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}