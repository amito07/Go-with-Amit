package main

import (
	"encoding/json"
	"fmt"
)

// capital Letter represents a Public
// small letter represents a Private
// In Go, if you want to export a variable or a function, you have to start with a capital letter
// If you want to make it private, you have to start with a small letter
// If you want to make it private to the package, you have to start with an underscore "_"

func main() {
	type course struct {
		Name     string `json:"name"`
		Price    int
		Platform string `json:"website"`
		Password string `json:"-"` //"-" means not returnable key
		Tags []string `json:"tags,omitempty"` // omitempty means if the value is nil, it will not return the key
	}

	courses := []course{
		{"Python Course", 199, "Udemy", "1234", []string{"python", "development", "programming"}},
		{"Django Course", 299, "Udemy", "1234", []string{"python", "development", "programming"}},
		{"React Course", 199, "Udemy", "1234", nil},
	}

	jsonResponse , err := json.MarshalIndent(courses, "", "\t")

	if err != nil{
		panic(err)
	}

	fmt.Printf("%s\n", jsonResponse)
}