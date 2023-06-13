package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := new(Person)
	p.Name, p.Age = "Tom", 18

	fmt.Println(p)
}
