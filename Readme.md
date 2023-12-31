**Example**

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Sskrill/in_memory_cache/cache"
)

// Example
func main() {
	db := cache.NewCache(time.Second * 5)
	db.Set(1, "Drake")
	result, ok := db.Get(1)
	if !ok {
		log.Fatal("not found")
	}
	fmt.Println(result)
	time.Sleep(time.Second * 6)
	result2, ok2 := db.Get(1)
	if !ok2 {
		log.Fatal("not found")
	}
	fmt.Println(result2)
}
```
**Result**
[result](![image](https://github.com/Sskrill/in_memory_cache_Sskrill/assets/154072620/5016e33e-6270-4107-a934-17e70128e76c))
