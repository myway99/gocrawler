package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"mytest04/crawler/gocrawler/engine"
	"errors"
)

func ItemSaver(
	index string) (chan engine.Item, error) {

	client, err := elastic.NewClient(
		//ElasticSearch server address and port
		elastic.SetURL("http://192.168.1.188:9200/"),
		// Must turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		return  nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: got item " +
				"#%d: %v", itemCount, item)
			itemCount++

			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error " +
					"saving item %v: %v",
						item, err)
			}
		}
	}()
	return out, nil
}

func Save(
	client *elastic.Client, index string,
	item engine.Item)  error {

	//Index():存储数据
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		// Index名
		Index(index).
		//type名（自动分配id）
		Type(item.Type).
		//Id(item.Id).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.
		Do(context.Background())

	//if err != nil {
	//	return err
	//}

	return err
	//fmt.Printf("%+v",resp)

}


