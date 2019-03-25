package parser

import (
	"regexp"
	"strings"
	"github.com/abnereel/crawler_dating_sites/crawler/engine"
	"github.com/abnereel/crawler_dating_sites/crawler/model"
)

var nameRe = regexp.MustCompile(`<div class="name nickname" >([\w\W]+)<img src="/public/home/images/mem_rz2.png"`)
var genderRe = regexp.MustCompile(`<li><span>性别：</span>([^<]+)</li>`)
var ageRe = regexp.MustCompile(`<li><span>年龄：</span>([^<]+)</li>`)
var heightRe = regexp.MustCompile(`<li><span>身高：</span>([^<]+)</li>`)
var incomeRe = regexp.MustCompile(`<li><span>年收入：</span>([^<]+)</li>`)
var marriageRe = regexp.MustCompile(`<li><span>婚况：</span>([^<]+)</li>`)
var educationRe = regexp.MustCompile(`<li><span>学历：</span>([^<]+)</li>`)
var workAddRe = regexp.MustCompile(`<li><span>工作地：</span>([^<]+)</li>`)

func ParseUser(contents []byte, url string) engine.ParseResult {
	user := model.User{}

	user.Name = strings.Trim(extractString(contents, nameRe), " ")
	user.Age = extractString(contents, ageRe)
	user.Gender = extractString(contents, genderRe)
	user.Height = extractString(contents, heightRe)
	user.Income = extractString(contents, incomeRe)
	user.Marriage = extractString(contents, marriageRe)
	user.Education = extractString(contents, educationRe)
	user.WorkAdd = extractString(contents, workAddRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url: url,
				Type: "hongniang",
				Id: "",
				Payload: user,
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

type UserParser struct {}

func (p *UserParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseUser(contents, url)
}

func (p *UserParser) Serialize() (name string, args interface{}) {
	return "ParserUser", nil
}

func NewUserParser() *UserParser {
	return &UserParser{}
}