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

	// array
	var arr1 [3]int
	var arr2 = [3]int{10, 20, 30}
	arr3 := [...]int{10, 20}
	fmt.Printf("%v %v %v\n", arr1, arr2, arr3)
	fmt.Printf("%v %v\n", len(arr3), cap(arr3))
	fmt.Printf("%T %T\n", arr2, arr3)

	// slice
	var sl1 []int
	sl2 := []int{}
	fmt.Printf("sl1: %[1]T %[1]v %v %v\n", sl1, len(sl1), cap(sl1))
	fmt.Printf("sl2: %[1]T %[1]v %v %v\n", sl2, len(sl2), cap(sl2))
	fmt.Println(sl1 == nil)
	fmt.Println(sl2 == nil)
	sl1 = append(sl1, 1, 2, 3)
	fmt.Printf("sl1: %[1]T %[1]v %v %v\n", sl1, len(sl1), cap(sl1))
	sl3 := []int{4, 5, 6}
	sl1 = append(sl1, sl3...)
	fmt.Printf("sl1: %[1]T %[1]v %v %v\n", sl1, len(sl1), cap(sl1))

	sl4 := make([]int, 0, 2)
	fmt.Printf("sl4: %[1]T %[1]v %v %v\n", sl4, len(sl4), cap(sl4))
	sl4 = append(sl4, 1, 2, 3, 4)
	fmt.Printf("sl4: %[1]T %[1]v %v %v\n", sl4, len(sl4), cap(sl4))
	sl5 := make([]int, 4, 6)
	fmt.Printf("sl5: %[1]T %[1]v %v %v\n", sl5, len(sl5), cap(sl5))
	sl6 := sl5[1:3]
	sl6[1] = 10
	fmt.Printf("sl5: %[1]T %[1]v %v %v\n", sl5, len(sl5), cap(sl5))
	fmt.Printf("sl6: %[1]T %[1]v %v %v\n", sl6, len(sl6), cap(sl6))
	sl6 = append(sl6, 2)
	fmt.Printf("sl5: %[1]T %[1]v %v %v\n", sl5, len(sl5), cap(sl5))
	fmt.Printf("sl6: %[1]T %[1]v %v %v\n", sl6, len(sl6), cap(sl6))
	sc6 := make([]int, len(sl5[1:3]))
	fmt.Printf("sl5 source of copy: %v %v %v\n", sl5, len(sl5), cap(sl5))
	fmt.Printf("sl6 dst copy before: %v %v %v\n", sc6, len(sc6), cap(sc6))
	copy(sc6, sl5[1:3])
	fmt.Printf("sl6 dst copy after: %v %v %v\n", sc6, len(sc6), cap(sc6))
	sc6[1] = 12
	fmt.Printf("sl5: %v %v %v\n", sl5, len(sl5), cap(sl5))
	fmt.Printf("sc6: %v %v %v\n", sc6, len(sc6), cap(sc6))

	sl5 = make([]int, 4, 6)
	fs6 := sl5[1:3:3]
	fmt.Printf("sl5: %v %v %v\n", sl5, len(sl5), cap(sl5))
	fmt.Printf("fs6: %v %v %v\n", fs6, len(fs6), cap(fs6))
	fs6[0] = 6
	fs6[1] = 7
	fs6 = append(fs6, 8)
	fmt.Printf("sl5: %v %v %v\n", sl5, len(sl5), cap(sl5))
	fmt.Printf("fs6: %v %v %v\n", fs6, len(fs6), cap(fs6))
	sl5[3] = 9
	fmt.Printf("sl5: %v %v %v\n", sl5, len(sl5), cap(sl5))
	fmt.Printf("fs6: %v %v %v\n", fs6, len(fs6), cap(fs6))
}
