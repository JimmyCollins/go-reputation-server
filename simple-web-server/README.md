# Simple Reputation Server

A simple web server used as an exercise in learning Go. Currently it just uses the VirusTotal API to look up the MD5 of a given SHA256 hash. Note that you need a valid API key for the VirusTotal API added in handlers.go - you can sign-up for VirusTotal for free.

### Dependencies

The following are required for the reputation server to work correctly.

The VirusTotal client for Go:

```
 go get github.com/VirusTotal/vt-go
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
