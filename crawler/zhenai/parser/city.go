package parser

import (
	"regexp"
	"github.com/abnereel/crawler_dating_sites/crawler/engine"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			Parser: NewProfileParser(),
		})
	}

	return result
}
