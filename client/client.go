package main

import (
	"context"
	"fmt"
	"log"

	pb "mocker/protocols/demo"

	"google.golang.org/protobuf/encoding/protojson"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// 连接到 gRPC 服务端
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := pb.NewDocInfoServiceClient(conn)

	// 调用服务端的 GetMessage 方法
	msg, err := client.GetDocDataInfo(context.Background(), &pb.GetDocDataInfoReq{Docid: "foo123"})
	if err != nil {
		log.Fatalf("Failed to get message: %v", err)
	}

	log.Printf("Received message: \n")
	buffer, err := protojson.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed protojson Marshal message: %v", err)
	}
	fmt.Println(string(buffer))
}
