# learn-grpc-go
 gRPC in Go

## What is gRPC?
gRPC is a high performance, open-source universal RPC framework. gRPC was initially developed by Google and is now a Cloud Native Computing Foundation (CNCF) project. It is based on the HTTP/2 protocol and uses Protocol Buffers as the interface description language.

## Why gRPC?
gRPC is a modern, open source, high-performance remote procedure call (RPC) framework that can run anywhere. It enables client and server applications to communicate transparently, and makes it easier to build connected systems. gRPC offers features such as authentication, load balancing, health checking, and observability.

## How does gRPC work?
gRPC uses Protocol Buffers as the interface description language. The client and server can be written in any of the supported languages, and gRPC will handle the rest. gRPC uses HTTP/2 for transport, and Protocol Buffers for serialization.

## What are the benefits of gRPC?
gRPC offers several benefits, including:
- Performance: gRPC is built on top of HTTP/2, which is faster and more efficient than HTTP/1.1.
- Interoperability: gRPC supports multiple programming languages, making it easy to build connected systems.
- Strong typing: gRPC uses Protocol Buffers for serialization, which provides strong typing and schema enforcement.
- Code generation: gRPC generates client and server code from a single interface definition file, making it easy to build and maintain services.

## What are the limitations of gRPC?
gRPC has some limitations, including:
- Complexity: gRPC can be more complex to set up and use than other RPC frameworks.
- HTTP/2: gRPC uses HTTP/2 for transport, which can be more difficult to debug than HTTP/1.1.
- Performance: While gRPC is fast, it may not be the best choice for all use cases.
- Scalability: auto scaling is challenging in gRPC.

## How to use gRPC in Go?
To use gRPC in Go, you need to:
1. Define a service interface using Protocol Buffers.
2. Implement the service interface in Go.
3. Generate client and server code using the `protoc` compiler.
4. Create a gRPC server and client in Go.
5. Start the gRPC server and make requests from the client.

## Pre-requisites
- https://grpc.io/docs/protoc-installation/
- Go
- Protocol Buffers (protoc) compiler `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest` this is use to generate Go code from .proto files
- gRPC Go plugin `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest` this is use to generate gRPC code in Go

## We need to create a simple gRPC **server** and **client** in Go

### Step 1: Define a service interface using Protocol Buffers
Create a new file called `stock.proto` and define a service interface for a simple stock service:

```proto
syntax = "proto3";

package stock;

service StockService {
	rpc GetStockPrice (StockRequest) returns (StockResponse) {}
}

message StockRequest {
	string symbol = 1;
}

message StockResponse {
	string symbol = 1;
	double price = 2;
	int64 timestamp = 3;
}
```

### Step 2: Implement the service interface in Go
Create a new file called `server.go` and implement the `StockService` interface:

```go
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "github.com/username/learn-grpc-go/stock"
)

type server struct {
	pb.UnimplementedStockServiceServer
}

func (s *server) GetStockPrice(ctx context.Context, req *pb.StockRequest) (*pb.StockResponse, error) {
	log.Printf("Received request for stock symbol: %s", req.GetSymbol())
	return &pb.StockResponse{
		Symbol:    req.GetSymbol(),
		Price:     100.0,
		Timestamp: 1630512000,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStockServiceServer(s, &server{})
	reflection.Register(s)

	log.Println("Starting gRPC server on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
```

### Step 3: Generate client and server code using the `protoc` compiler
Run the following command to generate the client and server code from the `stock.proto` file:

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative stock.proto
```

### Step 4: Create a gRPC client in Go
Create a new file called `client.go` and implement a simple gRPC client to call the `GetStockPrice` method:

```go
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "github.com/username/learn-grpc-go/stock"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewStockServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetStockPrice(ctx, &pb.StockRequest{Symbol: "AAPL"})
	if err != nil {
		log.Fatalf("Failed to get stock price: %v", err)
	}

	log.Printf("Received stock price for symbol %s: $%.2f", resp.GetSymbol(), resp.GetPrice())
}
```

### Step 5: Start the gRPC server and make requests from the client
Run the gRPC server by executing the `server.go` file:

```bash
go run server.go
```

Run the gRPC client by executing the `client.go` file:

```bash
go run client.go
```

You should see the following output:

```
2021/09/02 15:00:00 Starting gRPC server on port :50051
2021/09/02 15:00:00 Received request for stock symbol: AAPL
2021/09/02 15:00:00 Received stock price for symbol AAPL: $100.00
```

Congratulations! You have successfully created a simple gRPC server and client in Go.

## Conclusion
gRPC is a powerful and efficient RPC framework that can be used to build connected systems in Go. By following the steps outlined in this guide, you can create gRPC services that communicate transparently and efficiently. For more information on gRPC, check out the official [gRPC documentation](https://grpc.io/docs/).