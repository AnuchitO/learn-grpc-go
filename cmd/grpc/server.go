package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/anuchito/learn-grpc-go/flight"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var flights = []*flight.Flight{
	{AirlineCode: "AS", Number: "3567"},
	{AirlineCode: "DL", Number: "1234"},
	{AirlineCode: "AA", Number: "5678"},
}

type flightServer struct {
	flight.UnimplementedFlightsServer
}

func (s *flightServer) GetFlight(ctx context.Context, f *flight.Flight) (*flight.Flight, error) {
	fmt.Println("GetFlight")
	if f.AirlineCode == "AS" && f.Number == "3567" {
		return &flight.Flight{
			AirlineCode: "AS",
			Number:      "3567",
		}, nil
	}

	return &flight.Flight{}, nil
}

func (s *flightServer) GetFlightList(e *emptypb.Empty, stream flight.Flights_GetFlightListServer) error {
	fmt.Println("GetFlightList")
	for _, f := range flights {
		if err := stream.Send(f); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// step 1: create tcp listener on port 8080
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	// step 2: create grpc server
	// insecure connection
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	// step 3: register flight server
	server := &flightServer{}
	flight.RegisterFlightsServer(grpcServer, server)

	// step 4: start grpc server
	log.Println("starting grpc server on localhost:8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve:", err)
	}
}
