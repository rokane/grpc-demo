# gRPC Demo Application

This is an example application used to demonstrate an approach to structuring
gRPC servers such that they can handle multiple API versions.

Code is spoken about in more detail in the following blog post:

[Building APIs with gRPC and Go](https://medium.com/@ryan.okane8/building-apis-with-grpc-and-go-9a6d369d7ce)

The gRPC server itself contains two implementations of the `Greeter` service.
Both services contain the same two endpoints (as explained below), but with
different message types in an attempt to illustrate a `breaking change` leading
to a new version (`v2`).

## Build and Run

Build executable and run in terminal 1.

```sh
> go build ./cmd/greeter/...
> ./greeter
```

Call the api in terminal 2.

```sh
# API v1
> grpcurl -plaintext -v -d '{"name": "ryan"}' localhost:8080 greeter.v1.Greeter/SayHello
> grpcurl -plaintext -v -d '{"name": "ryan"}' localhost:8080 greeter.v1.Greeter/SayGoodbye

# API v2
> grpcurl -plaintext -v -d '{"first_name": "ryan", "last_name": "okane"}' localhost:8080 greeter.v2.Greeter/SayHello
> grpcurl -plaintext -v -d '{"first_name": "ryan", "last_name": "okane"}' localhost:8080 greeter.v2.Greeter/SayGoodbye
```
