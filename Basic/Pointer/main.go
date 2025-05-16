package main

import "fmt"

type Person struct {
	name string
	age  int
}

func changePerson(p *Person){
	p.name = "Tom"
	fmt.Println("Inside changePerson", *p)
}

func main() {

	p1 := Person{"John", 30}
	fmt.Println("Before updating...", p1)

	changePerson(&p1)
	fmt.Println("After updating...", p1)

}