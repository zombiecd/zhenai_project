package parse

import (
	"github.com/zombiecd/zhenai_project/engine"
	"regexp"
)

const city = `<a href="(http://album.zhenai.com/u/[0-9]+"[^>]*)">([^<]+)</a>`
func ParseCityUser(content []byte) engine.ParseResult {
	res := regexp.MustCompile(city)
	matcher := res.FindAllSubmatch(content,-1)
	result := engine.ParseResult{}

	for _,m := range matcher{
		name := string(m[2])
		result.Items=append(result.Items,"user"+name)
		result.Requests=append(result.Requests,engine.Request{
			Url: string(m[1]),
			ParseFunc: engine.NilParse,
		})
	}
	return result

}
