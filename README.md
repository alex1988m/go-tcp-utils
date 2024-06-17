# TCP Server and Handler Utilities

This repository provides a simple TCP server implementation in Go, along with various handlers for managing TCP connections. It includes a TCP server that listens for incoming connections and delegates the handling of these connections to different types of handlers.

## Packages

### `server`

The `server` package provides the core `TCPServer` struct and methods to start, stop, and serve TCP connections.

#### `TCPServer`

- **Fields:**
  - `ServerAddr`: Address on which the server listens.
  - `Handler`: An implementation of the `handler.TCPHandler` interface.
  - `l`: Internal listener for the server.

- **Methods:**
  - `Start() error`: Starts the TCP server.
  - `Stop() error`: Stops the TCP server.
  - `Serve() error`: Accepts incoming connections and delegates them to the handler.
  - `listen() error`: Sets up the TCP listener.

### `handler`

The `handler` package defines the `TCPHandler` interface and includes various implementations of this interface to handle TCP connections in different ways.

#### `TCPHandler` Interface

- **Methods:**
  - `Handle(client net.Conn) error`: Handles the incoming TCP connection.

#### `TCPControlHandler`

The `TCPControlHandler` executes a system command when a new connection is accepted.

- **Fields:**
  - `Command`: Command and arguments to be executed.

- **Methods:**
  - `Handle(client net.Conn) error`: Executes the specified command, directing the connection's input to the command's stdin and sending the command's stdout back to the connection.

#### `TCPProxyHandler`

The `TCPProxyHandler` forwards incoming connections to a specified target address.

- **Fields:**
  - `TargetAddr`: The address to which the connection should be proxied.

- **Methods:**
  - `Handle(client net.Conn) error`: Proxies the incoming connection to the target address, managing bidirectional data transfer between the client and the target.

## Usage

### Setting up the TCP Server

To create and start a TCP server, you need to specify the server address and the handler. For example:

```go
package main

import (
	"github/alex1988m/go-tcp-utils/handler"
	"github/alex1988m/go-tcp-utils/server"
	"log"
)

func main() {
	// Create a proxy handler
	proxyHandler := &handler.TCPProxyHandler{TargetAddr: "example.com:80"}

	// Create a TCP server with the proxy handler
	tcpServer := &server.TCPServer{
		ServerAddr: ":8080",
		Handler:    proxyHandler,
	}

	// Start the server
	if err := tcpServer.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer tcpServer.Stop()

	// Serve incoming connections
	if err := tcpServer.Serve(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
```


### Options

```
  -h, --help            help for go-tcp-utils
  -p, --proxy string    proxy host:port (default "localhost:4444")
  -t, --target string   target host:port (default "localhost:5555")
```

## go-tcp-utils server

base server command

### Options

```
  -h, --help   help for server
```

### Options inherited from parent commands

```
  -p, --proxy string    proxy host:port (default "localhost:4444")
  -t, --target string   target host:port (default "localhost:5555")
```

## go-tcp-utils server proxy

creates proxy server for the given tcp server

```
go-tcp-utils server proxy [flags]
```

### Options

```
  -h, --help   help for proxy
```

### Options inherited from parent commands

```
  -p, --proxy string    proxy host:port (default "localhost:4444")
  -t, --target string   target host:port (default "localhost:5555")
```

## go-tcp-utils server test-http

create test http server

```
go-tcp-utils server test-http [flags]
```

### Options

```
  -h, --help   help for test-http
```

### Options inherited from parent commands

```
  -p, --proxy string    proxy host:port (default "localhost:4444")
  -t, --target string   target host:port (default "localhost:5555")
```

## go-tcp-utils server control

start proxy tcp server which expose bash for client

```
go-tcp-utils server control [flags]
```

### Options

```
  -h, --help   help for control
```

### Options inherited from parent commands

```
  -p, --proxy string    proxy host:port (default "localhost:4444")
  -t, --target string   target host:port (default "localhost:5555")
```
