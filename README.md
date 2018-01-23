# GoKit
This repository contains useful code that we use in our Go projects.

## Random
Random generates cryptographically random strings. This can be used to generate variables such as session IDs.

## Web
### FileServer
FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at root. This implementation will always returns 404 Not Found if the request is a directory, and will not serve `index.html`.

### Server
Server extends the default HTTP server with graceful shutdown on receiving `SIGINT` or `SIGTERM`. The web server is a 1 to 1 replacement of http.Server's `ListenAndServe()`.

Example code:
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
