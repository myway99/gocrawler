package parser

import (
	"testing"
	"io/ioutil"
	"mytest04/crawler/model"
)

func TestParseProfile(t *testing.T) {

	contents, err := ioutil.ReadFile(
		"profile_test_data.html")

	if err != nil {
		panic(err)
	}

	//result := ParseProfile(contents, "风中的蒲公英")
	result := ParseProfile(contents, "安静的雪")
	//result := ParseProfile(contents)

	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 " +
			"element; but was %v ", result.Items)
	}

	profile := result.Items[0].(model.Profile)
	//profile.Name = "风中的蒲公英"
	profile.Name = "安静的雪"

	expected := model.Profile{
		//Age:        41,
		//Height:     158,
		//Weight:     48,
		//Income:     "3001-5000元",
		//Gender:     "女",
		//Name:       "风中的蒲公英",
		//Xinzuo:     "处女座",
		//Occupation: "公务员",
		//Marriage:   "离异",
		//House:      "已购房",
		//Hokou:      "四川阿坝",
		//Education:  "中专",
		//Car:        "未购车",

		Age:        34,
		Height:     162,
		Weight:     57,
		Income:     "3001-5000元",
		Gender:     "女",
		Name:       "安静的雪",
		Xinzuo:     "牡羊座",
		Occupation: "人事/行政",
		Marriage:   "离异",
		House:      "已购房",
		Hokou:      "山东菏泽",
		Education:  "大学本科",
		Car:        "未购车",

	}

	if profile != expected {
		t.Errorf("expected %v; but was %v",
			expected, profile)
	}

}


