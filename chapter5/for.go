package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	for x := 1; x <= 10; x++ {
		fmt.Println(x)
	}
}
