package config

const (
	// Parser names
	ParseUserList = "ParseUserList"
	ParseUser = "ParserUser"
	NilParser = "NilParser"

	// Service ports
	ItemSaverPort = 1234
	WorkerPort0 = 9000

	// ElasticSearch
	ElasticIndex = "dating_user"

	// RPC Endpoints
	ItemSaverRpc = "ItemServerService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	//Rate limiting
	Qps = 20
)