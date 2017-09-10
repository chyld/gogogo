package main

import (
	"fmt"
	"runtime"
	"sync"
)

type dog struct {
	name string
	age  int
}

var x int
var wg = sync.WaitGroup{}

func main() {
	threads := runtime.GOMAXPROCS(-1)
	fmt.Printf("threads: %#v\n", threads)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		cpu()
	}

	wg.Wait()

	// d := dog{name: "fido", age: 3}
	// da := &d
	// da.bark()
	// da.sit()
	// fmt.Printf("dog is %#v \n", d)
	// var a int32 = 3
	// fmt.Printf("a is %#v and type %T \n", a, a)
	// fmt.Printf("fib of 10 is %#v \n", fib(10))
	// fizzbuzz(20)
}

func cpu() {
	var x int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			for k := 0; k < 1000; k++ {
				x = i * j * k
			}
		}
	}
	fmt.Println(x)
	wg.Done()
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
