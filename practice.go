package main

func main() {
	var ch1 chan int

	// 受信専用
	// var ch2 <-chan int

	// 送信専用
	// var ch3 chan<- int

	ch1 = make(chan int)
	ch2 := make(chan int, 5)

	println(cap(ch1))
	println(cap(ch2))
}
