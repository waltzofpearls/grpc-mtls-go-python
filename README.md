# gRPC with Mutual TLS Between Go and Python

### Prereqs

Install Go support for Protocol Buffers and gRPC-Go:

```sh
# only on macOS
# other OS see: https://github.com/golang/protobuf#installation
brew install protobuf

# install proto file compiler plugin for Go
go get -u github.com/golang/protobuf/protoc-gen-go

# install these deps in a Go module folder (with go.mod)
go get -u github.com/golang/protobuf
go get -u google.golang.org/grpc
```

Create a new virtualenv project, and then install deps:

```sh
pip install grpcio grpcio-tools
```

### Run server and client

First things first, run `make` to generate Go and Python code from `api/metrics.proto` file,
and create all the certs needed when launching the server and client.

In one terminal window/split, run:

```sh
make server
```

In another terminal window/split, run:

```sh
make client
```

You will see gRPC response printed in the client side window/split.
