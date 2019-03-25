package parser

import (
	"regexp"
	"github.com/abnereel/crawler_dating_sites/crawler/engine"
	"github.com/abnereel/crawler_dating_sites/crawler/model"
)

var (
	nameRe      = regexp.MustCompile(`<th><a href="http://album.zhenai.com/u/[\d]+" target="_blank">([^<]+)</a></th>`)
	ageRe       = regexp.MustCompile(`<td width="180"><span class="grayL">年龄：</span>([0-9]+)</td>`)
	genderRe    = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
	heightRe    = regexp.MustCompile(`<td width="180"><span class="grayL">身&nbsp;&nbsp;&nbsp;高：</span>([^<]+)</td>`)
	marriageRe  = regexp.MustCompile(`<td width="180"><span class="grayL">婚况：</span>([^<]+)</td>`)
	educationRe = regexp.MustCompile(`<td><span class="grayL">学&nbsp;&nbsp;&nbsp;历：</span>([^<]+)</td>`)
	addressRe   = regexp.MustCompile(`<td><span class="grayL">居住地：</span>([^<]+)</td>`)
	idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
)


func parseProfile(contents []byte, url string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = extractString(contents, nameRe)
	profile.Age = extractString(contents, ageRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Height = extractString(contents, heightRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, educationRe)
	profile.Address = extractString(contents, addressRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url: url,
				Type: "zhenai",
				Id: extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParser struct {}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", nil
}

func NewProfileParser() *ProfileParser {
	return &ProfileParser{}
}
