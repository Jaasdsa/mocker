package main

import (
	"context"
	"log"
	"net"

	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		md, _ := metadata.FromIncomingContext(ctx)
		// 加入host传递给mocktool
		newMd := md.Copy()
		newMd["x-fress-mock-domain"] = []string{"test"}
		outCtx := metadata.NewOutgoingContext(ctx, newMd)

		cc, _ := grpc.DialContext(ctx, "localhost:50051", grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
		return outCtx, cc, nil
	}

	proxySrv := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()), // was previously needed for proxy to function.
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)),
	)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}
	proxySrv.Serve(lis)
}
