package main

import (
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"github.com/abnereel/crawler_dating_sites/crawler_distributed/config"
	"github.com/abnereel/crawler_dating_sites/crawler_distributed/persist"
	"github.com/abnereel/crawler_dating_sites/crawler_distributed/rpcsupport"
)
var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}

	return rpcsupport.ServerRpc(host, &persist.ItemServerService{
		Client: client,
		Index: index,
	})
}
