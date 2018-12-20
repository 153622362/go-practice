package parser

import (
	"io/ioutil"
	"testing"
	"../parser"
)

func TestParseCityList(t *testing.T) {
	//存下来文件，保证不会出错
	contents, err := ioutil.ReadFile(
		"citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := parser.ParseCityList(contents)
	const  resultSize  = 470
	expectedUrls := []string{ //期望链接
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}
	//expectedCities := []string{ //期望城市
	//	"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	//}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d" +
			"requests; but had %d", resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d:%s;but was %s", i, url, result.Requests[i].Url)
		}
	}

	//for i, city := range expectedCities {
	//	if result.Items[i].(string) != city {
	//		t.Errorf("expected city #%d:%s;but was %s", i, city, result.Items[i].(string))
	//	}
	//}
	//
	//if len(result.Items) != resultSize {
	//	t.Errorf("result should have %d" +
	//		"requests; but had %d", resultSize, len(result.Items))
	//}
}