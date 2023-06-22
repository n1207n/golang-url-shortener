# golang-url-shortener
A sample application written in Golang to implement URL shortener

## Setup
1. Create `.env` file by copying `.env_example` file
2. Make sure you have `redis-server` with version >= 6.0 installed and running
3. You can build the binary by `go build cmd/server.go`
   1. If you are targetting a destination path for executables, then make sure you copy .env file along with it.

## Disclaimers
I chose redis for faster prototyping instead of coming up with the schema and setting up a RBDMS.

For this implementation, the redis key expiration value is -1 therefore it is KEEPTTL, implying no expiration time. This is a good base to extend further for eviction policy. Redis does have a LFU eviction policy so you can configure your redis-server instance in such way. For the scalable volume, set a realistic TTL per redis key based on your data retention policy.