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
![result]([https://cdn.discordapp.com/attachments/592741750393536522/1190948796708950096/image.png?ex=65a3a87f&is=6591337f&hm=2333441261944f7bcdce7a439fc6dee052c65c89c687342ca79bc60c0699a6ca&](https://cdn.discordapp.com/attachments/592741750393536522/1190948978108407808/image.png?ex=65a3a8ab&is=659133ab&hm=a1cb46045cb05df7bbd11f7f99beb8055e52dd1d43381b74ac37a6a31191a878&)https://cdn.discordapp.com/attachments/592741750393536522/1190948978108407808/image.png?ex=65a3a8ab&is=659133ab&hm=a1cb46045cb05df7bbd11f7f99beb8055e52dd1d43381b74ac37a6a31191a878&)
