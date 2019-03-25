package persist

import (
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"github.com/abnereel/crawler_dating_sites/crawler/engine"
	"github.com/abnereel/crawler_dating_sites/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://www.hongniang.com/user/6/a/11011577.html",
		Type: "hongniang",
		Id:   "97046681",
		Payload: model.Profile{
			Name:      "安静的雪",
			Gender:    "女",
			Age:       "34岁",
			Height:    "162cm",
			Marriage:  "已婚",
			Education: "本科",
			Address:   "山东菏泽",
		},
	}

	//TODO: Try to start up elastic search here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"), //设置地址
		elastic.SetSniff(false),                  // Must turn off sniff in docker
	)
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
	resp, err := client.
		Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", *resp.Source)
	var actual engine.Item
	_ = json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	// Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
