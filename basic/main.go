package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	// scope
	"github.com/keisuke-333/go_practice/basic/alib"
	"github.com/keisuke-333/go_practice/basic/foo"
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

// pointer
func Double1(i *int) {
	*i = *i * 2
}

func Double2(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

// struct
type T struct {
	User
}

type User struct {
	Name string
	Age  int
}

type Users []*User

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func UpdateUser1(user User) {
	user.Name = "A"
	user.Age = 100
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 100
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u *User) SetName(name string) {
	u.Name = name
}

// interface
type Stringfy interface {
	ToString() string
}
type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
}

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

// test
func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
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

	// pointer
	var p1 int = 100
	fmt.Println(&p1)
	var p2 *int = &p1
	fmt.Println(p2)
	fmt.Println(*p2)
	*p2 = 300
	fmt.Println(p1)
	p1 = 200
	fmt.Println(*p2)
	Double1(&p1)
	fmt.Println(p1)
	Double1(p2)
	fmt.Println(*p2)
	var ps1 []int = []int{1, 2, 3}
	Double2(ps1)
	fmt.Println(ps1)

	// struct
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)
	user2 := User{Name: "user2", Age: 30}
	fmt.Println(user2)
	user3 := User{"user3", 40}
	fmt.Println(user3)
	user4 := new(User)
	fmt.Println(user4)
	user5 := &User{}
	fmt.Println(user5)
	UpdateUser1(user1)
	UpdateUser2(user5)
	fmt.Println(user1)
	fmt.Println(*user5)
	user6 := User{Name: "user6"}
	user6.SayName()
	user6.SetName("user6_test")
	user6.SayName()
	user_t := T{User: User{Name: "user7", Age: 30}}
	fmt.Println(user_t)
	fmt.Println(user_t.User)
	fmt.Println(user_t.User.Name)
	fmt.Println(user_t.Name)
	user_t.SetName("user8")
	fmt.Println(user_t.User)
	user9 := NewUser("user9", 33)
	fmt.Println(*user9)
	user10 := User{Name: "user10", Age: 40}
	user11 := User{Name: "user11", Age: 50}
	user12 := User{Name: "user12", Age: 60}
	user13 := User{Name: "user13", Age: 70}
	users := Users{}
	users = append(users, &user10, &user11, &user12, &user13)
	for _, u := range users {
		fmt.Println(*u)
	}
	mu1 := map[int]User{
		1: {Name: "user14", Age: 80},
		2: {Name: "user15", Age: 90},
	}
	for _, v := range mu1 {
		fmt.Println(v)
	}

	// interface
	infa1 := []Stringfy{
		&Person{Name: "Taro", Age: 20},
		&Car{Number: "123-456", Model: "AB-1234"},
	}
	for _, v := range infa1 {
		fmt.Println(v.ToString())
	}
	iferr1 := RaiseError()
	fmt.Println(iferr1.Error())
	iferr2, ok := iferr1.(*MyError)
	if ok {
		fmt.Println(iferr2.ErrCode)
	}
	ifp1 := &Point{100, "ABC"}
	fmt.Println(ifp1)

	// scope
	fmt.Println(foo.Max)
	fmt.Println(foo.ReturnMin())

	// test
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(0))
	tl1 := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(tl1))

	// os
	/*
		os.Exit(1)
		fmt.Println("Start")
	*/
	/*
		defer func() {
			fmt.Println("defer")
		}()
		os.Exit(0)
	*/
	osf1, _ := os.Create("foo.txt")
	osf1.Write([]byte("Hello\n"))
	osf1.WriteAt([]byte("Golang\n"), 6)
	osf1.Seek(0, os.SEEK_END)
	osf1.WriteString("test")

	osf2, err := os.Open("foo.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer osf2.Close()

	osbs := make([]byte, 128)
	osn, _ := osf2.Read(osbs)
	fmt.Println(osn)
	fmt.Println(string(osbs))

	// time
	ti1 := time.Now()
	fmt.Println(ti1.Zone())

	// math
	fmt.Println(math.Pi)
	fmt.Println(math.Max(1, 100))
	fmt.Println(math.Min(1, 100))
}
