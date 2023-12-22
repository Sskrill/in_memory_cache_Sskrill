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
