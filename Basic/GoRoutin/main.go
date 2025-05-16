package main

import (
	"fmt"
	"time"
)

func function1(name string){
	fmt.Println("Hello " + name)
}

func function2(num int64){
	var sum int64
	var i int64
	sum = 0
	for i = 0; i < num; i++ {
		sum += i
	}
	fmt.Println("Sum is: ", sum)
}

func function3(num1 int, num2 int){
	fmt.Println("Sum is: ", num1 + num2)
}

func main() {

	go function1("John")
	go function2(1000000000)
	go function3(10, 20)

	time.Sleep(2 * time.Second)
	fmt.Println("Main function")

}