package demo

import (
	"fmt"
	"time"
)

func createrWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {
	//var c chan int //c==nil
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createrWorker(i)

	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}
func worker(id int, c chan int) {
	for n := range c {
		//n,ok:=<-c
		//if !ok{
		//	break
		//}
		fmt.Printf("worker %d received %d\n", id, n)
	}
}
func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}
func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}
func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("bufferedChannel")
	bufferedChannel()
	fmt.Println("channelClose")
	channelClose()
}
