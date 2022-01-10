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

	ch2 <- 1
	println(len(ch2))

	ch2 <- 2
	ch2 <- 3
	println(len(ch2))

	i := <-ch2
	println(i)

	i2 := <-ch2
	println(i2)

	i3 := <-ch2
	println(i3)
}
