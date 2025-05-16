package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"name"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"` //"-" means not returnable key
	Tags []string `json:"tags,omitempty"` // omitempty means if the value is nil, it will not return the key
}

func main() {
	jsonFromWeb := [] byte(`
	        {
                "name": "Django Course",
                "Price": 299,
                "website": "Udemy",     
                "tags": ["python","development","programming"]
        	}
			`)
	var myCourse course
	
	checkValidity := json.Valid(jsonFromWeb)

	if checkValidity {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	}else{
		fmt.Println("JSON is not valid")
	}

	var myIncomingCourse map[string] interface{}

	json.Unmarshal(jsonFromWeb, &myIncomingCourse)

	fmt.Printf("%#v\n", myIncomingCourse)

	for key, value := range myIncomingCourse{
		fmt.Println("Key is: ", key, "Value is: ", value)
	}


}