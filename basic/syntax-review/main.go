package main

import (
	"fmt"
	"os"
	"syntax-review/calculator"

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
}
