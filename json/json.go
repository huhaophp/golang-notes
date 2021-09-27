package json

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", taginfo, " doc: ", tagdoc)
	}
}


func Demo() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"zhouxingchi", "zhangbozhi"}}

	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
	}

	fmt.Printf("%s\n", jsonStr)

	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
	}

	fmt.Printf("%v\n", myMovie)

	var u resume
	findTag(&u)
}
