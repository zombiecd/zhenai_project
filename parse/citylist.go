package parse

import (
	"github.com/zombiecd/zhenai_project/engine"
	"regexp"
)

const CityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`

//解析城市列表，得到Url和City
func ParseCity (content []byte) engine.ParseResult{
	res := regexp.MustCompile(CityListRe)
	matcher := res.FindAllSubmatch(content,-1)
	result := engine.ParseResult{}
	for _,m := range matcher{
	//	log.Printf("City:%v,Url:%v",string(m[2]),string(m[1]))
		result.Requests=append(result.Requests,engine.Request{
			Url: string(m[1]),
			ParseFunc: ParseCityUser,
		})
		result.Items=append(result.Items,"City:"+string(m[2]))
		//for _,item := range result.Requests{
		//log.Printf("go get item %v", item)
		//}
	}
	return result
}
