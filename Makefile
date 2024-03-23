greet:
	protoc -I ./greet/proto --go_opt=module=github.com/foxlovecat1989/golang-grpc-prac --go_out=./greet/proto --go-grpc_opt=module=github.com/foxlovecat1989/golang-grpc-prac --go-grpc_out=./greet/proto ./greet/proto/*.proto

.PHONY: greet

