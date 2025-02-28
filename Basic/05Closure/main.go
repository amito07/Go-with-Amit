package main

import "fmt"

const a = 100

var p int = 10

func outer() func(){
	money := 100
	age := 40

	fmt.Println("Age:", age)

	demo := func(){
		money = money + a + p
		fmt.Println("Money:", money)
	}
	return demo
}

func call(){
	inr1 := outer()
	inr1()
	inr1()

	inr2 := outer()
	inr2()
	inr2()
}


func main() {
	call()
}

func init() {
	fmt.Println("Init Function call first.....")
}