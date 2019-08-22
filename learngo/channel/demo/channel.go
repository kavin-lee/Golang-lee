package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		//go func() { //再开一个进行并行的运行
		//	done <- true
		//}()
		wg.Done()
	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func createrWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(id, w.in, wg)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	//var c chan int //c==nil
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createrWorker(i, &wg)

	}
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
		//<-workers[i].done
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-workers[i].done
	}
	//time.Sleep(time.Millisecond)
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
}
