package zhenai

import (
	"engine"
	"regexp"
)

const cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]+>([^<]+)</a>`

func CityList(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityListReg)
	s := string(contents)
	res := reg.FindAllStringSubmatch(s, -1)

	result := engine.ParseResult{}
	for _, m := range res {
		//放入城市名字到items
		result.Items = append(result.Items, m[2])
		//将每个城市对应的连接作为一个新的 Request 存起来
		result.Requests = append(result.Requests, engine.Request{
			Url:       m[1],
			ParseFunc: City,
		})
	}

	return result
}
