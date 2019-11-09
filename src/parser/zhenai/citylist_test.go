package zhenai

import (
	"io/ioutil"
	"testing"
)

func TestCityList(t *testing.T) {
	b, err := ioutil.ReadFile("zhenghun.html")
	if err != nil {
		panic(err)
	}

	//被测函数 检查返回结果是否符合预期
	r := CityList(b)

	const totalSize = 470
	exceptedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	exceptedCities := []string{
		"阿坝",
		"阿克苏",
		"阿拉善盟",
	}

	//测试:总数量是否符合预期
	if len(r.Requests) != totalSize {
		t.Errorf("total size is wrong, excepted: %d, actual：%d", totalSize, len(r.Requests))
	}

	//测试：城市数据
	for i, city := range exceptedCities {
		if city != r.Items[i].(string) {
			t.Errorf("city is wrong,excepted:%s, actual:%s", city, r.Items[i].(string))
		}
	}

	//测试：url数据
	for i, url := range exceptedUrls {
		if url != r.Requests[i].Url {
			t.Errorf("url is wrong,except ：%s, actual : %s", url, r.Requests[i].Url)
		}
	}
}
