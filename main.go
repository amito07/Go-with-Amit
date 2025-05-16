package main

import (
	"fmt"
)

// example 1
// instead of using any we can use interface{} but it is not type safe
// we can also use comparable types
func genericPrint[T comparable](list []T){
	for _, item := range list {
		println(item)
	}
}

// example 2
type stack[T comparable] struct{
	elements []T
}

// example 3
// use specific type
func genericItems[T int | string](list []T){
	for _, item := range list {
		println(item)
	}
}
func main(){
	//example 1
	genericPrint([]int{1, 2, 3})
	genericPrint([]string{"a", "b", "c"})

	//example 2
	// create a stack instance
	s := stack[int]{
		elements: []int{1,2,3},
	}

	s2 := stack[string]{
		elements: []string{"a", "b", "c"},
	}

	fmt.Println("Stack elements: ", s.elements)
	fmt.Println("Stack elements: ", s2.elements)

	//example 3
	genericItems([]int{1, 2, 3})
	genericItems([]string{"a", "b", "c"})
}