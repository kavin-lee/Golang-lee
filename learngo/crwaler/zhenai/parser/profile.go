package parser

import (
	"Golang-lee/learngo/crwaler/engine"
	"Golang-lee/learngo/crwaler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄: </span>([\d]+)岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况: </span>([^<]+)</td>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	//re := regexp.MustCompile(ageRe)
	match := ageRe.FindSubmatch(contents)

	if match != nil {
		age, err := strconv.Atoi(string(match[1]))
		if err != nil {
			profile.Age = age
		}
	}
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	profile.Marriage = extractString(contents, marriageRe)
	//re = regexp.MustCompile(marriageRe)
	//match = marriageRe.FindSubmatch(contents)
	//if match != nil {
	//	profile.Marriage = string(match[1])
	//}
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
