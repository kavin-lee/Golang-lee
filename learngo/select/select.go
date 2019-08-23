package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func worker(id int, c chan int) {
	for n := range c {
		//n,ok:=<-c
		//if !ok{
		//	break
		//}
		time.Sleep(1 * time.Second)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}
func createrWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func main() {
	var c1, c2 = generator(), generator()
	var worker = createrWorker(0)
	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	n := 0
	//hasValue := false
	for {
		var actieWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			actieWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			//fmt.Println("Received from c1:", n)
			//hasValue=true
			values = append(values, n)
		case n = <-c2:
			//fmt.Println("Received from c2:", n)
			//hasValue=true
			values = append(values, n)
			//default:
			//	fmt.Printf("No value received")
		case actieWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len=", len(values))
		case <-tm:
			fmt.Println("BYE")
			return
		}
	}

}
