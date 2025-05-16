package main

import (
	"fmt"
	"maps"
)

func main() {
	// Map is alterative to struct, dictionary, hash table
	// Map is a collection of key-value pairs

	// pre define map
	map1 := map[string]int{"price": 300, "quantity": 5}
	fmt.Println("Map", map1)

	// create a empty map
	map2 := make(map[string]int)
	fmt.Println("Map", map2)

	// add items to map
	map2["price"] = 400
	map2["quantity"] = 10
	map2["total"] = 4000
	fmt.Println("Map", map2)

	// delete item from map
	delete(map2, "price")
	fmt.Println("Map", map2)  //map[quantity:10 total:4000]

	// check if key is present in map
	value, is_present := map2["total"]
	if is_present {
		fmt.Println("Key is present", value)
	} else {
		fmt.Println("Key is not present")
	}

	// iterate over map
	for key, value := range map2 {
		fmt.Println("Key", key, "Value", value) //Key quantity Value 10
												// Key total Value 4000
	}

	// check if map is empty
	if len(map2) == 0 {
		fmt.Println("Map is empty")
	} else {
		fmt.Println("Map is not empty")
	}

	//Clear map
	clear(map2)
	fmt.Println("Map", map2) //map[]

	// copy map
	map3 := make(map[string]int)
	maps.Copy(map3, map1)
	map1["New Item"] = 100
	fmt.Println("Map1", map1) // map[New Item:100 price:300 quantity:5]
	fmt.Println("Map3", map3) // map[price:300 quantity:5]

	// compare map
	map4 := map[string]int{"price": 300, "quantity": 5}
	map5 := map[string]int{"quantity": 5, "price": 300}


	fmt.Println("Equal check", maps.Equal(map4, map5)) // true

}