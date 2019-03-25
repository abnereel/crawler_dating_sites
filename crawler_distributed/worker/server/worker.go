package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/abnereel/crawler_dating_sites/crawler_distributed/rpcsupport"
	"github.com/abnereel/crawler_dating_sites/crawler_distributed/worker"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
	}
	log.Fatal(
		rpcsupport.ServerRpc(
			fmt.Sprintf(":%d", *port),
			worker.CrawlService{}))
}
