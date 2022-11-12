package main

import (
	"fmt"
	"strconv"
)

// variable length arguments
func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

// channel
func reciever1(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

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

	// type switches
	var ts1 interface{} = 3
	switch v := ts1.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println(v, "I don't know")
	}

	// slice capacity
	sc1 := make([]int, 5, 10)
	fmt.Println(len(sc1))
	fmt.Println(cap(sc1))
	sc1 = append(sc1, 1, 2, 3, 4, 5, 6)
	fmt.Println(len(sc1))
	fmt.Println(cap(sc1))

	// slice copy
	sco1 := []int{1, 2, 3, 4, 5}
	sco2 := make([]int, 5, 10)
	sco3 := copy(sco2, sco1)
	fmt.Println(sco3, sco2)

	// variable length arguments
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	vla1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Sum(vla1...))

	// channel
	ch1 := make(chan int, 5)
	fmt.Println(cap(ch1))
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	fmt.Println(len(ch1))
	ch2 := <-ch1
	fmt.Println(ch2)
	fmt.Println(len(ch1))
	fmt.Println(<-ch1)
	fmt.Println(len(ch1))
	// channel & goroutine
	ch3 := make(chan int)
	ch4 := make(chan int)
	go reciever1(ch3)
	go reciever1(ch4)
	chi1 := 1
	for chi1 < 30 {
		ch3 <- chi1
		ch4 <- chi1
		chi1++
	}
	// close
	chc1 := make(chan int, 2)
	chc1 <- 1
	close(chc1)
	chc2, ok := <-chc1
	fmt.Println(chc2, ok)
	chc3, ok := <-chc1
	fmt.Println(chc3, ok)
	chcf := make(chan int, 3)
	chcf <- 1
	chcf <- 2
	chcf <- 3
	close(chcf)
	for i := range chcf {
		fmt.Println(i)
	}
}
