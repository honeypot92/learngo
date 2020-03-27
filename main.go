package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"doe", "eho"}
	hj := person{name: "hyunjun", age: 29, favFood: favFood}
	fmt.Println(hj)
}
