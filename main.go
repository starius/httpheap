package main

import(
    "fmt"
)

func main() {
    c1 := make(chan int)
    c2 := make(chan int)
    go func() {
        c1 <- 0
        c1 <- 10
        c1 <- 20
        close(c1)
    }()
    go func() {
        c2 <- 5
        c2 <- 15
        c2 <- 25
        close(c2)
    }()

    chan_heap := new(ChanHeap)
    chan_heap.AddChan(c1)
    chan_heap.AddChan(c2)
    for {
        value, has_value := chan_heap.PopValue()
        if !has_value {
            break
        }
        fmt.Println(value)
    }
}
