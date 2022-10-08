package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World")

	// byte
	b1 := []byte("a")
	b2 := []byte("あ")
	b3 := string(b1)
	b4 := string(b2)
	fmt.Println(b1, b2, b3, b4)
	b5 := []byte{227, 129, 130}
	fmt.Println(string(b5))

	// anonymous function
	af1 := func(x, y int) int {
		return x * y
	}(3, 6)
	fmt.Println(af1)

	// error handling
	e1 := "100"
	e2, err := strconv.Atoi(e1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(e2)

	// for range
	fr1 := map[string]int{"apple": 100, "banan": 200, "orange": 300}
	for k, v := range fr1 {
		fmt.Println(k, v)
	}
}
