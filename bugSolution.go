```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        data := make(map[int]int)
        var mu sync.Mutex
        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        mu.Lock()
                        data[i] = i * 2
                        mu.Unlock()
                }(i)
        }
        wg.Wait()
        fmt.Println(len(data))
}
```