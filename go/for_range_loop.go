package main

import "fmt"

func main() {
    array := [4]int{0,1,2,3}
    
    for i, v := range array {
        fmt.Printf("[%v]: %v\n", i, v)
    }
}