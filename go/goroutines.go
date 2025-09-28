package main

import (
	"fmt"
	"sync"
)

func goroutine(name string, waitgroup *sync.WaitGroup) {
	fmt.Println(name)
	waitgroup.Done()
}

func main() {
	var waitgroup sync.WaitGroup
	names := []string{"Dominic", "Ashley", "Nora", "Zack"}

	for _, name := range names {
		waitgroup.Add(1)
		go goroutine(name, &waitgroup)
	}

	waitgroup.Wait()
}
