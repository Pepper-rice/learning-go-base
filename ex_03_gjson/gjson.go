package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	jsonStr := `
		{
		  "name": {"first": "Tom", "last": "Anderson"},
		  "age":37,
		  "children": ["Sara","Alex","Jack"],
		  "fav.movie": "Deer Hunter",
		  "friends": [
			{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
			{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
			{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
		  ]
		}
	`

	fmt.Println(gjson.Get(jsonStr, "name"))
}
