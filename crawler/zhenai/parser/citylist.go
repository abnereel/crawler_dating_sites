package parser

import (
	"regexp"
	"github.com/abnereel/crawler_dating_sites/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return result
}
