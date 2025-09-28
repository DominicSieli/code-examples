package main

import (
	"fmt"
	"sync"
)

func mutateInt(intChannel chan<- int, buffer int, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()

	for i := 0; i < buffer; i++ {
		intChannel <- i
	}

	close(intChannel)
}

func printInt(intChannel <-chan int, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()

	for v := range intChannel {
		fmt.Printf("%d ", v)
	}

	fmt.Printf("\n")
}

func main() {
	const buffer int = 20
	var waitgroup sync.WaitGroup
	intChannel := make(chan int)

	waitgroup.Add(2)
	go mutateInt(intChannel, buffer, &waitgroup)
	go printInt(intChannel, &waitgroup)

	waitgroup.Wait()
}
