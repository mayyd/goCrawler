package engine

import (
	"fetcher"
	"fmt"
	"log"
)

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func Nilparser([]byte) ParseResult {
	return ParseResult{}
}

func Run(seeds ...Request) {
	//将所有要请求的放进队列等待处理
	var requestQueue []Request
	for _, seed := range seeds {
		requestQueue = append(requestQueue, seed)
	}

	for len(requestQueue) > 0 {
		r := requestQueue[0]
		requestQueue = requestQueue[1:]

		log.Printf("Fetching:%s\n", r.Url)

		//抓回结果
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetch url : %s ,have error : %s ", r.Url, err.Error())
			continue
		}

		//解析结果
		parseRes := r.ParseFunc(body)

		//将解析后的新的请求放入队列中
		requestQueue = append(requestQueue, parseRes.Requests...)

		for k, item := range parseRes.Items {
			fmt.Println("Got item :", item, ",url:", parseRes.Requests[k].Url)
		}
	}
}
