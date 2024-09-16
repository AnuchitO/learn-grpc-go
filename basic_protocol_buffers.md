# Steps
1. create package `flight` then create `flight.proto` file
1. use "github.com/golang/protobuf/proto" package to Marshal and Unmarshal the message
1. store the message in map[string][]byte to simulate a database and key is generate
1. add and get the message from the `map` using the key as the ID

# Steps gRPC
The Go code generator does not produce output for services by default. To generate service definitions, you need to use the `protoc` compiler with the `protoc-gen-go-grpc` plugin.
[Service](https://protobuf.dev/reference/go/go-generated/#service)

1. create a new file called `flight.proto` and define a service interface for a simple flight service:
```proto
syntax = "proto3";

package flight;

service FlightService {
    rpc AddFlight (Flight) returns (FlightID) {}
    rpc GetFlight (FlightID) returns (Flight) {}
}

message Flight {
    string id = 1;
    string airline = 2;
    string number = 3;
    string departure = 4;
    string arrival = 5;
}

message FlightID {
    string id = 1;
}
```

2. Implement the service interface in Go
Create a new file called `server.go` and implement the `FlightService` interface:
```go
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "github.com/username/learn-grpc-go/flight"
)

type server struct {
	pb.UnimplementedFlightServiceServer
}

func (s *server) AddFlight(ctx context.Context, in *pb.Flight) (*pb.FlightID, error) {
	log.Printf("Received: %v", in)
	return &pb.FlightID{Id: in.Id}, nil
}

func (s *server) GetFlight(ctx context.Context, in *pb.FlightID) (*pb.Flight, error) {
	log.Printf("Received: %v", in)
	return &pb.Flight{
		Id: in.Id,
		Airline: "Airline",
		Number: "Number",
		Departure: "Departure",
		Arrival: "Arrival",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFlightServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

3. Create a new file called `client.go` and implement the client to call the `FlightService` interface:
```go
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "github.com/username/learn-grpc-go/flight"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFlightServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddFlight(ctx, &pb.Flight{
		Id: "1",
		Airline: "Airline",
		Number: "Number",
		Departure: "Departure",
		Arrival: "Arrival",
	})
	if err != nil {
		log.Fatalf("could not add flight: %v", err)
	}
	log.Printf("Flight ID: %s", r.Id)

	r2, err := c.GetFlight(ctx, &pb.FlightID{Id: "1"})
	if err != nil {
		log.Fatalf("could not get flight: %v", err)
	}
	log.Printf("Flight: %v", r2)
}
```
