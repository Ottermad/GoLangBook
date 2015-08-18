package main

import "fmt"

// var x string = "Hello World" // Will not cause an error

func main() {
	var x string = "Hello World" // Will cause an error
	fmt.Println(x)
}

func f() {
	fmt.Println(x)
}