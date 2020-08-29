//go:generate protoc -I api/proto --proto_path=api/proto/greeter/v1alpha1 --go_out=pkg/api --go_opt=paths=source_relative api/proto/greeter/v1alpha1/greeter_api.proto
//go:generate protoc -I api/proto --proto_path=api/proto/greeter/v1alpha2 --go_out=pkg/api --go_opt=paths=source_relative api/proto/greeter/v1alpha2/greeter_api.proto
