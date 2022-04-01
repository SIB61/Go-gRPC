Server:
	go run "/home/sabit/GoLang/Go-gRPC/server/."

Client:
	go run "/home/sabit/GoLang/Go-gRPC/client/."

Gen:
	protoc --go_out=.  protos/user.proto
	protoc --go-grpc_out=.  protos/user.proto