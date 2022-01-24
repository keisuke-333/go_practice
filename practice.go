package main

import "fmt"

func main() {
	// c := make([]int, 0, 5)
	c := make([]int, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}
