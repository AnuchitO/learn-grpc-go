# schema
- `xxx.proto` file

compile proto file
```bash
protoc --go_out=. xxx.proto
```

will generate `xxx.pb.go` file


# Contents
- Install `protobuf` compiler and Go plugins(`protoc-gen-go`, `protoc-gen-go-grpc`)
- Define a **Messages** in a `.proto` file test by compile to generate Go code file
-




# Steps
1. install Protocol Buffers (protoc) compiler `brew install protobuf`
  - check version `protoc --version`
1. install plugin  plugin to generate Go code `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - the `protoc` compiler will use this plugin to generate Go code from .proto files
1. install gRPC Go plugin `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
  - the `protoc` compiler will use this plugin to generate gRPC code in Go


1. create package `flight` then create `flight.proto` file
```proto
syntax = "proto3";
```
1. run `protoc --go_out=. flight.proto` compile to generate Go code file
  - will show error

  ```
  protoc-gen-go: unable to determine Go import path for "flight.proto"

  Please specify either:
          • a "go_package" option in the .proto source file, or
          • a "M" argument on the command line.
  ```

  [docs package](https://protobuf.dev/reference/go/go-generated/#package)

1. update `flight.proto` file
```proto
syntax = "proto3";

option go_package = "../flight"; // This is the package name that will be used in the generated Go code

package flight;
```

1. run `protoc --go_out=. flight.proto` compile to generate Go code file should be happy now

1. create message Flight
```proto
syntax = "proto3";

option go_package = "../flight"; // This is the package name that will be used in the generated Go code

package flight;

message Flight {
    string airlineCode = 1;
    string number = 2;
}
```

1. run `protoc --go_out=. flight.proto` compile to generate Go code file should be generate `flight.pb.go` file with `Flight` struct

1. define `message Flight { string number = 1; } `
  - `<field_name> <data_type> = <tag_number>;`
  - `field_name`: The name of the field.
  - `data_type`: The type of data the field can hold. In this case, all fields are strings (string) except for price, which is a double.
    - https://protobuf.dev/programming-guides/proto3/#scalar
  - `tag_number`: A unique integer that identifies the field within the message. This number is used for serialization and deserialization.
1. run `protoc --go_out=. flight.proto` compile to go file
  - will show error

  [docs package](https://protobuf.dev/reference/go/go-generated/#package)

  - In order to generate Go code, the Go package’s import path must be provided for every .proto file
  - There are two ways to provide the Go package’s import path:
    1. add `option go_package = "github.com/anuchito/learn-grpc-go/flight";` to `flight.proto` or
    1. use `--go_opt` flag with the `M` argument ` protoc --go_opt=Mflight.proto=./ --go_out=. flight.proto`
      - The Go import path may be specified on the command line when invoking the compiler, by passing one or more `M${PROTO_FILE}=${GO_IMPORT_PATH}` flags. Example usage:
      `${GO_IMPORT_PATH}` is relative to the current directory
  - We recommend declaring it within the .proto file

## define gRPC service
1. define `message FlightList { repeated Flight flights = 1; }`
  - `repeated Flight flights = 1;` is a repeated field, which is similar to a list or array in other programming languages.
1. define `service FlightService { rpc GetFlights(Flight) returns (FlightList); }`
1. run `protoc --go_out=. --go-grpc_out=. flight.proto` compile to generate Go code file should be generate `flight_grpc.pb.go` file with `FlightServiceServer` interface
  - Reference: https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code

## note
- `xx.proto` add top file `option go_package = "google.golang.org/grpc/examples/route_guide/routeguide";`
- `protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC`
- `-I` flag to specify the directory where the .proto file is located
- protobuf syntax version 3 guide https://protobuf.dev/programming-guides/proto3/

- difference between field `string airlineCode = 1;` vs `string airline_name = 2;`
  - `AirlineCode` is a json tag in Go struct will be `json:"airlineCode"`
  - `airline_code` is a field tag in proto file will be `json:"airline_code"`

- protoc-gen-go: invalid Go import path "flight" for "flight.proto"
  - The import path must contain at least one period ('.') or forward slash ('/') character.
  - See https://protobuf.dev/reference/go/go-generated#package for more information.

- when `option go_package = ".";` will generate file in the same directory as the proto file but the package name will be `package __`
- hacking `option go_package = "../flight";` will generate file in the parent directory of the proto file but the package name will be `package flight`
- gRPC example :https://github.com/grpc/grpc-go/tree/master/examples/helloworld
