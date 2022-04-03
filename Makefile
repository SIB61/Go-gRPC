Server:
	go run "/home/sabit/GoLang/Go-gRPC/server/."

Client:
	go run "/home/sabit/GoLang/Go-gRPC/client/."

Pb:
	protoc --go_out=.  protos/user.proto
	protoc --go-grpc_out=.  protos/user.proto