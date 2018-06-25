package persist

import (
	"context"
	"encoding/json"
	"mytest04/crawler/gocrawler/engine"
	"mytest04/crawler/gocrawler/model"
	"testing"

	"gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		//ElasticSearch server address and port
		elastic.SetURL("http://192.168.1.188:9200/"),
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// Save expected item
	err = Save(client, index, expected)

	if err != nil {
		panic(err)
	}

	// Fetch saved item
	// 反序列化：json解析
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", *resp.Source)

	var actual engine.Item
	//json.Unmarshal([]byte(resp.Source), &actual)
	json.Unmarshal(*resp.Source, &actual)

	//if err != nil {
	//	panic(err)
	//}

	actualProfile, _ := model.FromJsonObj(
		actual.Payload)
	actual.Payload = actualProfile

	// Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v",
			actual, expected)
	}

}
