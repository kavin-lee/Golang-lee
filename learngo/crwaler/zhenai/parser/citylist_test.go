package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile(
		"citylist_test_data.html")
	if err != nil {
		panic(err)

	}
	//fmt.Printf("%s\n",contents)22
	result := ParseCityList(contents)
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d"+
			"results:but had %d", resultSize,
			len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d"+
			"results:but had %d", resultSize,
			len(result.Items))
	}
}
