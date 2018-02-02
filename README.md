# GoKit
[![GoDoc](https://godoc.org/github.com/moexmen/gokit?status.svg)](https://godoc.org/github.com/moexmen/gokit)

This repository contains useful code that we use in our Go projects.

## Examples

### web.Server
```
import (
	"log"
	"time"

	"github.com/moexmen/gokit/web"
)

func main() {
	s := web.Server{
		Addr:    ":8080",
		Timeout: 5 * time.Second,
	}
	log.Println("Starting...")
	log.Println(s.ListenAndServe())
}
```
If you use docker, [docker stop](https://docs.docker.com/compose/reference/stop/) has a default timeout of 10 seconds, the graceful timeout should be set to expire before then.
