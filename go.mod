module mocker

go 1.20

require (
	github.com/jhump/protoreflect v1.16.0
	github.com/mwitkow/grpc-proxy v0.0.0-20220126150247-db34e7bfee32
	github.com/mwitkow/grpc-proxy-v2023 v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.61.0
	google.golang.org/protobuf v1.33.1-0.20240408130810-98873a205002

)

require (
	github.com/bufbuild/protocompile v0.10.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231106174013-bbf56f31fb17 // indirect
)

// replace github.com/mwitkow/grpc-proxy-v2022 => github.com/mwitkow/grpc-proxy v0.0.0-20220126150247-db34e7bfee32

replace github.com/mwitkow/grpc-proxy-v2023 => github.com/mwitkow/grpc-proxy v0.0.0-20230212185441-f345521cb9c9
