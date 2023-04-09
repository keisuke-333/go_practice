package main

import (
	"fmt"
	"os"
	"syntax-review/calculator"
	"unsafe"

	"github.com/joho/godotenv"
)

const (
	Mac int = iota + 1
	Windows
	Linux
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))
	fmt.Println(calculator.Offset)
	fmt.Println(calculator.Sum(1, 2))
	fmt.Println(calculator.Multiply(1, 2))

	// variables
	var i1 int = 1
	var i2 = 2
	i3 := 3
	fmt.Printf("i1: %v %T\n", i1, i1)
	fmt.Printf("i2: %v %T\n", i2, i2)
	fmt.Printf("i3: %v %T\n", i3, i3)

	var f1 float32 = 1.1
	f2 := 1.2
	fmt.Printf("f1: %[1]v %[1]T, f2: %[2]v %[2]T\n", f1, f2)

	val1, val2 := 1.3, "go"
	fmt.Println(val1, val2)

	val3 := 1
	val4 := 1.2
	val5 := float64(val3) + val4
	fmt.Printf("val: %v, type: %T\n", val5, val5)

	fmt.Printf("Mac: %v, Windows: %v, Linux: %v\n", Mac, Windows, Linux)

	// pointer
	var ui1 uint16
	fmt.Printf("memory address of ui1: %p\n", &ui1)
	var p1 *uint16
	fmt.Printf("value of p1: %v\n", p1)
	p1 = &ui1
	fmt.Printf("value of p1: %v\n", p1)
	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1))
	fmt.Printf("memory address of p1: %p\n", &p1)
	fmt.Printf("value of ui1(dereference) %v\n", *p1)
	*p1 = 1
	fmt.Printf("value of ui1: %v\n", ui1)

	var pp1 **uint16 = &p1
	fmt.Printf("value of pp1: %v\n", pp1)
	fmt.Printf("value of p1(dereference) %v\n", *pp1)
	fmt.Printf("value of ui1(dereference) %v\n", **pp1)
	**pp1 = 10
	fmt.Printf("value of ui1: %v\n", ui1)

	// shadowing
	ok, result := true, "A"
	fmt.Printf("memory address of result: %p\n", &result)
	if ok {
		result := "B"
		fmt.Printf("memory address of result: %p\n", &result)
		println(result)
	} else {
		result := "C"
		println(result)
	}
	println(result)
}
