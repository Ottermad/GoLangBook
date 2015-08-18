package main

import "fmt"

func main() {
	const x string = "Hello World" 
	x = "Hello Again" // This will error
	fmt.Println(x)
}