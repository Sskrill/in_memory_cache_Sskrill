My first proj "in_memory_cache"

This project can help you with data caching.

Example
```go
package main

import (
	"fmt"

	"github.com/Sskrill/in_memory_cache/cache"
)

// Example
func main() {

	cache := cache.NewMemoryCache()

	cache.Set("Oleg", 45)

	Oleg := cache.Get("Oleg")

	fmt.Println(Oleg)

	cache.Delete("Vass")

	Oleg = cache.Get("Vass")
	fmt.Println(Oleg)

}
```
