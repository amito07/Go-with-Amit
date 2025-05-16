package main

import "fmt"

func main() {
	// type 1
	fmt.Println("Type 1")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("================================================")

	// type 2
	fmt.Println("Type 2")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("================================================")

	// type 3
	fmt.Println("Type 3")
	for {
		fmt.Println("Infinite Loop")
		break
	}

	fmt.Println("================================================")

	for n := range []int{0, 1, 2, 3, 4, 5} {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}