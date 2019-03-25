package parser

/*
import (
	"regexp"
	"strconv"
	"github.com/abnereel/crawler_dating_sites/crawler/engine"
	"github.com/abnereel/crawler_dating_sites/crawler/model"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([0-9]+)岁</div>`)
var genderRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">(\d+)cm</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">(\d+)</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var occupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var hokouRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		profile.Weight = weight
	}
	profile.Gender = extractString(contents, genderRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
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
*/
