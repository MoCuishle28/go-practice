package parser

import(
	"testing"
	"io/ioutil"
)


// 测试
func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d, result.Requests had %d", resultSize, len(result.Requests))
	}
}