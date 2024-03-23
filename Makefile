greet:
	protoc -I ./greet/proto --go_opt=module=github.com/foxlovecat1989/golang-grpc-prac --go_out=./greet/proto --go-grpc_opt=module=github.com/foxlovecat1989/golang-grpc-prac --go-grpc_out=./greet/proto ./greet/proto/*.proto
	go build -o bin/greet/server ./greet/server
	go build -o bin/greet/client ./greet/client

greet-server-run:
	./bin/greet/server

greet-client-run:
	./bin/greet/client

.PHONY: greet

