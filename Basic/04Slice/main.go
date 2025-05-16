package main

import (
	"fmt"
	"slices"
)

func main() {
	var slice1 = make([]int, 0, 5)
	// add elements to the slice
	slice1 = append(slice1, 1)
	slice1 = append(slice1, 2)
	slice1 = append(slice1, 3)
	slice1 = append(slice1, 4)
	slice1 = append(slice1, 5)

	// add more elements to the slice
	// capacity is doubled when the slice is full
	// fmt.Println(cap(slice1))
	// fmt.Println(len(slice1))
	// fmt.Println(slice1)

	// copy the slice to another slice
	// slice2 := make([]int, len(slice1))
	// copy(slice2, slice1)
	// fmt.Println("========================== Brefore ========================")
	// fmt.Println("Slice 1", slice1)
	// fmt.Println("Slice 1 Capacity", cap(slice1))
	// fmt.Println("Slice 1 Length", len(slice1))
	// fmt.Println("Slice 2",slice2)
	// fmt.Println("Slice 2 Capacity", cap(slice2))
	// fmt.Println("Slice 2 Length", len(slice2))

	// fmt.Println("========================== After (add items in Slice 1) ========================")
	// slice1 = append(slice1, 6)
	// fmt.Println("Slice 1 Capacity", cap(slice1))
	// fmt.Println("Slice 1 Length", len(slice1))
	// fmt.Println("Slice 2 Capacity", cap(slice2))
	// fmt.Println("Slice 2 Length", len(slice2))

	// fmt.Println("=========================== After (add items in Slice 2) ========================")
	// slice2 = append(slice2, 7)

	// fmt.Println("Slice 1 Capacity", cap(slice1))
	// fmt.Println("Slice 1 Length", len(slice1))
	// fmt.Println("Slice 2 Capacity", cap(slice2))
	// fmt.Println("Slice 2 Length", len(slice2))

	// Slice Operator
	slice3 := slice1[0:2]
	fmt.Println("Slice 3", slice3)

	// Slice method
	slice4 := []int{0,1,2,3,5,6,7,8,10}
	// slice5 := []int{1,2}

	// fmt.Println(slices.Compare(slice4, slice5))
	index, is_present := slices.BinarySearch(slice4, 5)
	fmt.Println("Index", index)
	fmt.Println("Is Present", is_present)

}