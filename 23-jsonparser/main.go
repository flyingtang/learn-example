package main

import (
	"github.com/buger/jsonparser"
	"fmt"
)

func main(){

	data := []byte(`{
  "person": {
    "name": {
      "first": "Leonid",
      "last": "Bugaev",
      "fullName": "Leonid Bugaev"
    },
    "github": {
      "handle": "buger",
      "followers": 109
    },
    "avatars": [
      { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
    ]
  },
  "company": {
    "name": "Acme"
  }
}`)

	fn, _ := jsonparser.GetString(data, "person", "name", "fullName")
	fmt.Println(fn)

	// There is `GetInt` and `GetBoolean` helpers if you exactly know key data type
	jsonparser.GetInt(data, "person", "github", "followers")

	// When you try to get object, it will return you []byte slice pointer to data containing it
	// In `company` it will be `{"name": "Acme"}`
	jsonparser.Get(data, "company")


}