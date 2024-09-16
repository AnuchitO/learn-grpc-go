package main

import (
	"context"
	"io"
	"log"

	"github.com/anuchito/learn-grpc-go/flight"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	// step 1: create grpc client
	serverURL := "localhost:8080"
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(serverURL, opts...)
	if err != nil {
		log.Fatal("failed to connect:", err)
	}
	defer conn.Close()

	
	// step 2: create flight client
	client := flight.NewFlightsClient(conn)

	f, err := client.GetFlight(context.Background(), &flight.Flight{})
	if err != nil {
		log.Fatal("failed to get flight:", err)
	}
	log.Printf("get flight success: %#v\n", f)

	stream, err := client.GetFlightList(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatal("failed to get flight list:", err)
	}

	for {
		f, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.ListFeatures failed: %v", err)
		}
		log.Printf("steam flight list success: %#v\n", f)
	}

}
