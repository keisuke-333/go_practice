package main

import "fmt"

func main() {
	f := 1.11
	fmt.Printf("%T\n", f)
	fmt.Printf("%T %v\n", int(f), int(f))

	// [5 6]
	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[2:4])

	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v", m, m)
}
