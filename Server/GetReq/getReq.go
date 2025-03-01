package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	demoUrl := "https://jsonplaceholder.typicode.com/posts"

	content, err := http.Get(demoUrl)

	if err != nil{
		panic(err)
	}

	defer content.Body.Close()

	// fmt.Println("Status Code: ", content.Status)
	// fmt.Println("Content Length: ", content.ContentLength)
	data, err := ioutil.ReadAll(content.Body)

	if err != nil{
		panic(err)
	}
	//=========================== One way =====================================

	// One way we can convert bytes to strings

	// fmt.Println("Content in binary format", string(data))



	//=============================== Another way =================================
	// Another way we can convert bytes to strings using "strings" package

	var responseData strings.Builder

	byteCount, _ := responseData.Write(data)

	fmt.Println("Bytes count: ", byteCount)

	fmt.Println("Content are", responseData.String())

}