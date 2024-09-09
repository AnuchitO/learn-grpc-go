.PHONY: compile
compile:
	protoc --go_out=. flight.proto

.PHONY: compilev2
compilev2:
	protoc -I . flight.proto --go_out=.

.PHONY: grpc
grpc:
	 protoc --go_out=. --go-grpc_out=. flight.proto
