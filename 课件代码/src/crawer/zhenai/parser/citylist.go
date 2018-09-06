package parser

import (
	"crawer/engine"
	"regexp"
)
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"
			[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	//exp := `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"
	//		[^>]*>([^<]+)</a>`
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents,-1)

	result :=engine.ParseResult{}
	for _, m:= range matchs{
		result.Items = append(result.Items,string( m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:string(m[1]),
			PaserFunc:engine.NilParser,
		})
		//fmt.Printf("City: %s, URL %s\n",m[2],m[1])
	}
	//fmt.Printf("Count: %d\n",len(matchs))
	return result
}
