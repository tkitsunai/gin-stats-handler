[golang-stats-api-handler](https://github.com/fukata/golang-stats-api-handler) available for framework "gin"

## Install

```
go get github.com/tkitsunai/gin-stats-handler
```

or,

If you use "dep" for package manager.

```
dep ensure -add github.com/tkitsunai/gin-stats-handler
```

If you use "glide" for package manager.

```
glide get github.com/tkitsunai/gin-stats-handler
```


## Usage Example
```
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tkitsunai/gin-stats-handler"
)

func main() {
	g := gin.Default()
	g.GET("/stats", gin-stats-handler.Stats())
	g.Run()
}
```

## Response

```
{
	time: 1533935033459278600,
	go_version: "go1.10.3",
	go_os: "darwin",
	go_arch: "amd64",
	cpu_num: 4,
	goroutine_num: 4,
	gomaxprocs: 4,
	cgo_call_num: 1,
	memory_alloc: 972552,
	memory_total_alloc: 972552,
	memory_sys: 5343480,
	memory_lookups: 6,
	memory_mallocs: 8485,
	memory_frees: 280,
	memory_stack: 360448,
	heap_alloc: 972552,
	heap_sys: 2785280,
	heap_idle: 786432,
	heap_inuse: 1998848,
	heap_released: 0,
	heap_objects: 8205,
	gc_next: 4473924,
	gc_last: 0,
	gc_num: 0,
	gc_per_second: 0,
	gc_pause_per_second: 0,
	gc_pause: [ ]
}
```

## using in framework "gin-gonic"

- [Gin](https://github.com/gin-gonic/gin)
