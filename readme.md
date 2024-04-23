# 生成带描述符的 pb

protoc --go_out=. --descriptor_set_out=demo.pb demo.proto

# 生成带有 service 的桩代码

protoc --go_out=protocols/demo --go_opt=paths=source_relative --go-grpc_out=protocols/demo --go-grpc_opt=paths=source_relative --descriptor_set_out=demo.pb demo.proto

# 1. 运行 grpc mocker 服务端代码

```bash
go run .
```

# 2. 运行 grpc mocker 客户端代码

```bash
go run client/client.go
```

- mocker 服务端已经实现动态类型 mocker
- 不用每次生成桩代码，大大提高了灵活性
