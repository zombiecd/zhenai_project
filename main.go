package main

import (
	"github.com/zombiecd/zhenai_project/engine"
	"github.com/zombiecd/zhenai_project/parse"
)

func main() {

	engine.Run(engine.Request{
		Url:       "https://www.zhenai.com/zhenghun",
		ParseFunc: parse.ParseCity,
	})

	//读取页面信息并打印
	//all, err := fetch.Fetch("https://www.zhenai.com/zhenghun")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s\n", all)

	//e := parse.ParseCity(all)
	//fmt.Printf("%v",e.Requests)
}
