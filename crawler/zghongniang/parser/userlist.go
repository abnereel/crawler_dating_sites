package parser

import (
	"regexp"
	"github.com/abnereel/crawler_dating_sites/crawler/engine"
)

const userListRe = `<div class="name"> <a target="_blank" href="(.+)">(.+)</a> <img src="/public/home/images/mem_rz2.png"`

func ParseUserList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(userListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        "http://www.hongniang.com" + string(m[1]),
			Parser: NewUserParser(),
		})
	}

	return result
}
