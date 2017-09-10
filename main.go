package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ready!")
	nums := []int8{3, 5, 7, 9}
	c := make(chan int8)

	go add(nums[:2], c)
	go add(nums[2:], c)

	a, b := <-c, <-c
	fmt.Printf("%#v\n", a+b)
}

func add(s []int8, c chan int8) {
	var sum int8
	for _, v := range s {
		sum += v
	}
	c <- sum
}
