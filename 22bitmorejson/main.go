package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON with more details")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 399, "learncodeonline.in", "bcd123", []string{"fullstack", "js"}},
		{"Angular Bootcamp", 199, "learncodeonline.in", "cde123", nil},
	}

	// package this data as json data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
            "coursename": "ReactJS Bootcamp",
            "website": "learncodeonline.in",
            "tags": ["web-dev","js"]
    	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	// some cases where we want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v, value is %v and Type is %T\n", k, v, v)
	}
}
