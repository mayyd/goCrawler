package zhenai

import (
	"engine"
	"regexp"
)

//匹配每个用户的访问地址
const userReg = `<a href="(http://album.zhenai.com/u/\d+)" target="_blank">([^<]+)</a>`

func City(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(userReg)
	res := reg.FindAllStringSubmatch(string(contents), -1)

	r := engine.ParseResult{}
	for _, v := range res {
		r.Items = append(r.Items, "User: "+string(v[2]))
		r.Requests = append(r.Requests, engine.Request{
			Url:       v[1],
			ParseFunc: engine.Nilparser,
		})
	}

	return r
}
