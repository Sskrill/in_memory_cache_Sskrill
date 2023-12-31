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
![result]([![image](https://github.com/Sskrill/in_memory_cache_Sskrill/assets/154072620/5016e33e-6270-4107-a934-17e70128e76c)](https://cdn.discordapp.com/attachments/592741750393536522/1190948978108407808/image.png?ex=65a3a8ab&is=659133ab&hm=a1cb46045cb05df7bbd11f7f99beb8055e52dd1d43381b74ac37a6a31191a878&)https://cdn.discordapp.com/attachments/592741750393536522/1190948978108407808/image.png?ex=65a3a8ab&is=659133ab&hm=a1cb46045cb05df7bbd11f7f99beb8055e52dd1d43381b74ac37a6a31191a878&)
