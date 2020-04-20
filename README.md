# Simple Reputation Server

![Go](https://github.com/JimmyCollins/go-reputation-server/workflows/Go/badge.svg)

A simple web server that serves reputations of files, built from the ground up as an exercise in learning Go. Currently it uses MongoDB as the backend data store, and assumes it is running on the local machine.

### Dependencies

The following are required for the reputation server to work correctly.

The MongoDB Go libraries:

```
 go.mongodb.org/mongo-driver/mongo
```

Mux for routing:

```
 go get github.com/gorilla/mux
```

Zizuo for caching:

```
 go get github.com/arriqaaq/zizou
```

### Building

```
 go build
```

### Examples

Get details for a file, given a SHA256 hash of a file:

```
 http://localhost:8080/file/ab7da8511c7698ab10eaf29e0f5597b1616de9325be5124f72fb9eed26a6750e
```

Insert a new reputation in MongoDB:

```
 http://localhost:8080/insert/ab7da8511c7698ab10eaf29e0f5597b1616de9325be5124f72fb9eed26a6750e/bad
