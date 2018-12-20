package parser

import (
	"../../engine"
	"regexp"
)
//const cityRe  = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
var (
		profile1Re  = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	   //cityUrlRe = regexp.MustCompile(
	   //	`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func parseCity(contents []byte) engine.ParseResult {
	matches := profile1Re.FindAllSubmatch(contents , -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		//result.Items = append(
		//	result.Items, "User " + name)
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: 				url,
			ParserFunc : func(c []byte) engine.ParseResult {
				return ParseProfile(c, url, name)
			},
		})
	}

	//matches = cityUrlRe.FindAllSubmatch(contents, -1)
	//
	//for _, m := range matches {
	//	result.Requests = append(result.Requests, engine.Request{
	//		Url:	string(m[1]),
	//		ParserFunc:parseCity,
	//	})
	//}
	return result
}