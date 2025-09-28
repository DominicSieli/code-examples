package main

import "fmt"

func main() {
	cache := make(map[int]string)

	cache[0] = "zero"
	cache[1] = "one"
	cache[2] = "two"
	cache[3] = "three"
	delete(cache, 3)

	for i := 0; i < 4; i++ {
		fmt.Printf("cache[%d]: %s\n", i, cache[i])
	}
}
