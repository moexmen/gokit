# GoKit
Useful code for our Go projects

`go get -u github.com/moexmen/gokit`

## Server
Starts a HTTP server with graceful shutdown when receiving `SIGINT` or `SIGTERM`

[Docker Stop](https://docs.docker.com/compose/reference/stop/) has a default timeout of 10 seconds. We recommend to use a timeout lower than 10 seconds if you use Docker.

### Credits
Credits to ArdanLab for providing an awesome base code for HTTP server shutdown https://github.com/ardanlabs/gotraining/tree/master/topics/web/shutdown
