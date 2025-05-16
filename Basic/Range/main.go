package main

import "fmt"

func main() {
	// Range is used to iterate over a collection of elements
	// It can be used with arrays, slices, maps, strings, and channels
	// It returns two values: the index and the value
	// The index is the position of the element in the collection
	// The value is the element itself

	// Example with array
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println("Index:", i, "Value:", v)
	}

	// Example with slice
	slice := []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		fmt.Println("Index:", i, "Value:", v)
	}

	// Example with map
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v)
	}

	// Example with string
	// Note: In Go, strings are UTF-8 encoded, so the range loop will iterate over runes
	// which are Unicode code points. This means that if the string contains multi-byte characters,
	// the index may not correspond to the byte index of the string.
	// For example, the string "hello" has 5 bytes and 5 runes
	
	str := "hello"
	for i, v := range str {
		fmt.Println("Index:", i, "Value:", string(v))
	}
}