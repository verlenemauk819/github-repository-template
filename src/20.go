package main

import "fmt"

func main() {
    // Example usage of goroutines and channels
    go func() {
        fmt.Println("Hello from goroutine")
    }()
    
    var x int = 0
    
    for i := 0; i < 5; i++ {
        channel := make(chan int, 1)
        
        go func(channel chan int) {
            for n := range channel {
                x += n
            }
            
            fmt.Println("Sum:", x)
        }(x)
    }

    fmt.Println("The sum is:", x)
}
