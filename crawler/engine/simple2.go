package engine

import (
	"log"
	"strconv"
)

func (e SimpleEngine) Run2(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	//获取用户列表
	r := requests[0]
	requests = requests[1:]
	page := 1
	isPage := true
	for {
		if isPage {
			r.Url = r.Url + strconv.Itoa(page)
			page++
		} else {
			r = requests[0]
			requests = requests[1:]
		}
		parseResult, err := Worker(r)
		if err != nil {
			continue
		}
		if len(parseResult.Requests) == 0 {
			isPage = false
			break
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %s", item)
		}
		/*if (page == 2) {
			break
		}*/
	}
}
