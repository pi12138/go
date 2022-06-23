package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello, world")

	// define variable
	// style 1
	var age int = 10
	var name string = "张三"
	var str string = "hello, world"

	// style 2
	var weigth = 100

	// style 3
	height := 180

	// variable must be used
	fmt.Println(age, name, str, weigth, height)

	// define const
	// const dont use
	const PI = 3.1415926
	const ONE int = 1
}
