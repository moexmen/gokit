# GoKit
Useful code for our Go projects

`go get -u github.com/moexmen/gokit`

## Random
Cryptographically random strings to be used for things like session IDs.

## Server
Starts a HTTP server with graceful shutdown when receiving `SIGINT` or `SIGTERM`

[Docker Stop](https://docs.docker.com/compose/reference/stop/) has a default timeout of 10 seconds. We recommend to use a timeout lower than 10 seconds if you use Docker.
