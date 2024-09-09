package main

import (
	"fmt"
	"log"

	"github.com/anuchito/learn-grpc-go/flight"
	"google.golang.org/protobuf/proto"
)

func main() {

	msg := &flight.Flight{
		AirlineCode: "AS",
		Number:      "3567",
	}
	fmt.Printf("flight struct: %#v\n", msg)

	b, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("marshal error:", err)
	}
	fmt.Printf("byetes: %b\n", b)

	fmt.Println()
	var flt flight.Flight
	if err := proto.Unmarshal(b, &flt); err != nil {
		log.Fatal("unmarshal error:", err)
	}
	fmt.Printf("unmarshaled struct: %#v\n", &flt)
}
