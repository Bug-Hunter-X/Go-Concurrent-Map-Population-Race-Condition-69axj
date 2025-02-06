```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        data := make(map[int]int)
        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        data[i] = i * 2
                }(i)
        }
        wg.Wait()
        fmt.Println(len(data))
}
```