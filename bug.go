```go
package main

import "fmt"

func main() {
    var m map[string]int
    m["a"] = 1 // This will panic if m is not initialized properly.
    fmt.Println(m["a"])
}
```