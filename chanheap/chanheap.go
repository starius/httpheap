package chanheap

import(
    "container/heap"
)

type item struct {
    channel <-chan int
    last_value int
}

type ChanHeap []item

func (h ChanHeap) Len() int {
    return len(h)
}

func (h ChanHeap) Less(i, j int) bool {
    return h[i].last_value < h[j].last_value
}

func (h ChanHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *ChanHeap) Push(x interface{}) {
    *h = append(*h, x.(item))
}

func (h *ChanHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func (h *ChanHeap) AddChan(channel <-chan int) {
    value, has_value := <-channel
    if has_value {
        heap.Push(h, item{channel, value})
    }
}

func (h *ChanHeap) PopValue() (int, bool) {
    if len(*h) == 0 {
        return 0, false
    }
    first_element := (*h)[0]
    value := first_element.last_value
    new_value, has_new_value := <-first_element.channel
    if has_new_value {
        (*h)[0].last_value = new_value
        heap.Fix(h, 0)
    } else {
        heap.Remove(h, 0)
    }
    return value, true
}
