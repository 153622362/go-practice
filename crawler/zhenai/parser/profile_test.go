package parser

import (
	"../../model"
	"../parser"
	"io/ioutil"
	"testing"
	"../../engine"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile(
		"profile_test_data.html")

	if err != nil {
		panic(err)
	}

	name := "宁静致远"

	result := parser.ParseProfile(contents,"http://album.zhenai.com/u/1972182204", string(name))
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 " +
			"element;but was %v", result.Items)
	}

	//profile := result.Items[0].(model.Profile)
	actual := result.Items[0]
	expected := engine.Item{
		Url : "http://album.zhenai.com/u/1972182204",
		Type: "zhenai",
		Id: "1972182204",
		Payload:model.Profile{
			Name:		"宁静致远",
			Age:		43,
			Height:		161,
			Income:		"5001-8000元",
			Marriage:	"离异",
			Education:	"大学本科",
			Hokou:		"阿坝",
		},
	}

	if actual != expected {
		t.Errorf("expected %v but was %v", expected, actual)
	}
}