package main

import(
    "fmt"
    "net/http"

    "github.com/starius/httpheap/chanheap"
)

func ReadHttpInts(url string, channel chan<- int) {
    defer close(channel)
    res, err := http.Get(url)
    if err != nil {
        return
    }
    defer res.Body.Close()
    for {
        var value int
        _, err := fmt.Fscanf(res.Body, "%d\n", &value)
        if err != nil {
            break
        }
        channel <- value
    }
}

func main() {
    const N = 10
    chans := make([]chan int, 0)
    for step := 1; step <= N; step++ {
        url := fmt.Sprintf(
            "http://127.0.0.1:25516/ints/?start=1&stop=100&step=%d",
            step,
        )
        channel := make(chan int)
        chans = append(chans, channel)
        go ReadHttpInts(url, channel)
    }

    chan_heap := new(chanheap.ChanHeap)
    for _, channel := range chans {
        chan_heap.AddChan(channel)
    }
    for {
        value, has_value := chan_heap.PopValue()
        if !has_value {
            break
        }
        fmt.Println(value)
    }
}
