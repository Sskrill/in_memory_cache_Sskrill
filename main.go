package main

import (
	"fmt"
	"time"

	"github.com/Sskrill/in_memory_cache/cache"
)

// Example
func main() {

	cache := cache.NewMemoryCache()

	cache.Set("Oleg", 45, time.Second*5)

	Oleg := cache.Get("Oleg")

	fmt.Println(Oleg)
	time.Sleep(time.Second * 6) //We imagined that we had a delay

	Oleg = cache.Get("Oleg") // exit code (not found info )
	fmt.Println(Oleg)

}
