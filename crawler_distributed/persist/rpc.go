package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"log"
	"github.com/abnereel/crawler_dating_sites/crawler/engine"
	"github.com/abnereel/crawler_dating_sites/crawler/persist"
)

type ItemServerService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemServerService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}
	return err
}
