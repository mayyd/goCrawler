package main

import (
	"engine"
	"parser/zhenai"
)

func main() {
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: zhenai.CityList,
	})
}
