package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func testReg() {
	s := `
My email is 8899666@qq.com
email1 is abc@def.org
email2 is    kkk@qq.com
email3 is ddd@abc.com.cn
`
	reg, _ := regexp.Compile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+\.[a-zA-Z0-9.]+)`)

	r := reg.FindAllStringSubmatch(s, -1)
	for i := range r {
		fmt.Printf("res = %s @ %s", r[i][1], r[i][2])
		fmt.Println()
	}
}

func printCityList(contents []byte) {
	//fmt.Println(string(contents))
	//<a target="_blank" href="http://www.zhenai.com/zhenghun/beijing" data-v-7e67c21c>北京</a>
	reg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]+>([^<]+)</a>`)
	s := string(contents)
	res := reg.FindAllStringSubmatch(s, -1)
	for _, m := range res {
		fmt.Println(m[2], m[1])
	}
}

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//检查返回状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("return status code is not ok:", resp.StatusCode)
		return
	}

	//防止乱码，因此使用第三方包来处理编码问题
	e := DeterminEncoding(resp.Body)
	//编码转换
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	s, _ := ioutil.ReadAll(utf8Reader)

	printCityList(s)
}

func DeterminEncoding(r io.Reader) encoding.Encoding {
	b, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(b, "")
	return e
}
