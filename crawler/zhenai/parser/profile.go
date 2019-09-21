package parser

import (
	"go-practice/crawler/engine"
	"go-practice/crawler/model"
	"regexp"
	"strconv"
	"strings"
)

//var ageRe  = regexp.MustCompile(`<div[^>]+>[^>]+ ([0-9]+)岁 [^<]+</div>`)
//var marriageRe = regexp.MustCompile(`<div[^>]+>海西 | 33岁 | 大学本科 | 离异 | 160cm | 3001-5000元</div>`)

//var heightRe = regexp.MustCompile(`<div[^>]+>海西 | 33岁 | 大学本科 | 离异 | 160cm | 3001-5000元</div>`)
//var incomeRe = regexp.MustCompile(`<div[^>]+>[^>]+ [0-9\-]+元</div>`)
//var educationRe = regexp.MustCompile(`<div[^>]+>海西 | 33岁 | 大学本科 | 离异 | 160cm | 3001-5000元</div>`)

//var occupationRe = regexp.MustCompile(`<li><span>职业：</span>[^<]+</li>`)
//var hokouRe = regexp.MustCompile(`<div[^>]+>[\w]+\s\|[^<]+</div>`)
//var xinzuoRe = regexp.MustCompile(`<li><span>星座：</span>[^<]+</li>`)
//var houseRe = regexp.MustCompile(`<li><span>是否购房：</span>[^<]+</li>`)
//var carRe = regexp.MustCompile(`<li><span>是否买车：</span>[^<]+</li>`)
var numberRe = regexp.MustCompile("\\w{2,3}")

var profileRe = regexp.MustCompile(`<div[^>]+>([^>]+ [0-9]+岁 [^<]+)</div>`)

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, url string,
	name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profilec := profileRe.FindSubmatch(contents)
	if len(profilec) < 1 {
		return engine.ParseResult{}
	}
	s := strings.Split(string(profilec[1]), "|")

	//fmt.Println(s[1])
	//fmt.Println("sss")
	age, err := strconv.Atoi(string(numberRe.FindSubmatch([]byte(s[1]))[0]))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(string(numberRe.FindSubmatch([]byte(s[4]))[0]))
	if err == nil {
		profile.Height = height
	}
	profile.Marriage = strings.Replace(s[3], " ", "", -1)
	profile.Income = strings.Replace(s[5], " ", "", -1)
	profile.Education = strings.Replace(s[2], " ", "", -1)
	profile.Hokou = strings.Replace(s[0], " ", "", -1)
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai", //elastirsearch 表名
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}
	return result
}

func extractString(
	contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
