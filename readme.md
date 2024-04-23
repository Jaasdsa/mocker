# 生成带描述符的 pb

protoc --go_out=. --descriptor_set_out=demo.pb demo.proto

# 生成带有 service 的桩代码

protoc --go_out=protocols/demo --go_opt=paths=source_relative --go-grpc_out=protocols/demo --go-grpc_opt=paths=source_relative --descriptor_set_out=demo.pb demo.proto
