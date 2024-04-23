package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"encoding/base64"
	"io/ioutil"

	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 自定义handleStream处理请求和响应
	server := grpc.NewServer(grpc.UnknownServiceHandler(handleStream))

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func handleStream(srv interface{}, stream grpc.ServerStream) error {
	// 获取请求流的目标接口名称
	fullMethodName, ok := grpc.MethodFromServerStream(stream)
	// fullMethodName =  /demo.DocInfoService/GetDocDataInfo
	if !ok {
		return status.Errorf(codes.Internal, "lowLevelServerStream not exists in context")
	}
	log.Printf("fullMethodName: %s", fullMethodName)

	// 获取请求流的元数据
	md, _ := metadata.FromIncomingContext(stream.Context())
	log.Printf("md: %s", md)

	// 读取编译后的 protobuf 文件
	fileBytes, err := ioutil.ReadFile("demo.pb")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	fileDescSet := new(descriptorpb.FileDescriptorSet)
	if err := proto.Unmarshal(fileBytes, fileDescSet); err != nil {
		log.Fatalf("failed to unmarshal descriptor set: %v", err)
	}

	// 序列化 FileDescriptorSet 对象为字节数组
	bytes, err := proto.Marshal(fileDescSet)
	if err != nil {
		log.Fatalf("failed to marshal descriptor set: %v", err)
	}

	// 转换字节数组为 Base64 字符串
	base64Str := base64.StdEncoding.EncodeToString(bytes)
	// fmt.Println(base64Str)

	// 解码 Base64 字符串为字节数组
	decodedBytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		log.Fatalf("failed to decode base64 string: %v", err)
	}

	// 反序列化字节数组为 FileDescriptorSet 对象
	decodedFileDescSet := new(descriptorpb.FileDescriptorSet)
	if err := proto.Unmarshal(decodedBytes, decodedFileDescSet); err != nil {
		log.Fatalf("failed to unmarshal descriptor set: %v", err)
	}

	// 使用 protoreflect 库解析文件描述符
	fileDesc, err := desc.CreateFileDescriptorFromSet(decodedFileDescSet)
	if err != nil {
		log.Fatalf("failed to create file descriptor: %v", err)
	}

	// 从文件描述符中查找 消息的描述符
	pos := strings.LastIndex(fullMethodName, "/")
	if pos < 0 {
		log.Fatalf("failed to cstrings.LastIndex(fullMethodName),r: %v", err)
	}
	svcName := fullMethodName[:pos]
	if pos := strings.Index(svcName, "/"); pos != -1 {
		svcName = svcName[pos+1:]
	}
	methodName := fullMethodName[pos+1:]
	svc := fileDesc.FindService(svcName)
	method := svc.FindMethodByName(methodName)
	res := dynamic.NewMessage(method.GetOutputType())

	// 根据 json 数据填入动态类型中
	// mockResStr := `{"data": {"tid": 256}}`
	reader := strings.NewReader(mockResStr)
	unmarshaler := jsonpb.Unmarshaler{}
	if err := unmarshaler.Unmarshal(reader, res); err != nil {
		log.Fatalf("Failed to unmarshal json data: %v", err)
	}

	if err := stream.SendMsg(res); err != nil {
		return status.Errorf(codes.Internal, "failed to send message: %s", err)
	}
	return nil
}
