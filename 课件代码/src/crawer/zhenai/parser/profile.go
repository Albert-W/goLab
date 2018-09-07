package parser

import (
	"crawer/engine"
	"crawer/model"
	"regexp"
	"strconv"
)

const ageRe = `<td><span class="label">年龄：</span>(\d+)岁</td>`
var ageRex = regexp.MustCompile(ageRe)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^>]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^>]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^>]+)</span></td>`)

const marriageRe = `<td><span class="label">婚况：</span>([^>]+)</td>`
var marriageRex = regexp.MustCompile(marriageRe)

var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^>]+)</td>`)
var OccupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^>]+)</td>`)
var HokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^>]+)</td>`)
var HouseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^>]+)</span></td>`)
var CarRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^>]+)</span></td>`)
//var nameRe = regexp.MustCompile(`<td><span class="label">姓名：</span>([^>]+)</td>`)




func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	//re := regexp.MustCompile(ageRe)
	//match := ageRex.FindSubmatch(contents)
	//
	//if match != nil {
	//	age, e := strconv.Atoi(string(match[1]))
	//	if e != nil{
	//		profile.Age = age
	//	}
	//}
	profile.Name = name

	age, e := strconv.Atoi(extractString(contents, ageRex))
	if e == nil {
		profile.Age = age
	}
	height, e := strconv.Atoi(extractString(contents, heightRe))
	if e == nil {
		profile.Height = height
	}
	weight, e := strconv.Atoi(extractString(contents, weightRe))
	if e == nil {
		profile.Weight = weight
	}

	//re = regexp.MustCompile(marriageRe)
	//match = marriageRex.FindSubmatch(contents)
	//
	//if match != nil {
	//	profile.Marriage = string(match[1])
	//}
	profile.Income = extractString(contents,incomeRe)
	profile.Gender = extractString(contents,genderRe)
	profile.Car = extractString(contents,CarRe)
	profile.Xinzuo = extractString(contents,xinzuoRe)
	profile.Education = extractString(contents,educationRe)
	profile.Occupation = extractString(contents,OccupationRe)
	profile.Hokou = extractString(contents,HokouRe)
	profile.House = extractString(contents,HouseRe)
	//profile.Name = extractString(contents,nameRe)
	profile.Marriage = extractString(contents,marriageRex)

	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string  {
	match := re.FindSubmatch(contents)
	if len(match) >=2 {
		return string(match[1])
	} else {
		return ""
	}

}
