package main

import (
	"fmt"
)

type dog struct {
	name string
	age  int
}

var x int

func main() {
	d := dog{name: "fido", age: 3}
	da := &d
	da.bark()
	da.sit()
	fmt.Printf("dog is %#v \n", d)
	var a int32 = 3
	fmt.Printf("a is %#v and type %T \n", a, a)
	fmt.Printf("fib of 10 is %#v \n", fib(10))
	fizzbuzz(20)
}

func (d dog) bark() {
	fmt.Printf("%#v is barking \n", d)
}

func (d *dog) sit() {
	fmt.Printf("%#v is sitting \n", d)
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func fizzbuzz(n int) {
	for ; n >= 0; n-- {
		if n%15 == 0 {
			fmt.Printf("fizzbuzz: %#v \n", n)
		} else if n%5 == 0 {
			fmt.Printf("fizz: %#v \n", n)
		} else if n%3 == 0 {
			fmt.Printf("buzz: %#v \n", n)
		}
	}
}
