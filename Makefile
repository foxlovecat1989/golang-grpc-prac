greet-proto-gen:
	protoc -I ./greet/proto --go_opt=module=github.com/foxlovecat1989/golang-grpc-prac --go_out=./greet/proto --go-grpc_opt=module=github.com/foxlovecat1989/golang-grpc-prac --go-grpc_out=./greet/proto ./greet/proto/*.proto

greet-build:
	go build -o bin/greet/server ./greet/server
	go build -o bin/greet/client ./greet/client

greet-server-run:
	./bin/greet/server

greet-client-run:
	./bin/greet/client


calculator-proto-gen:
	protoc -I ./calculator/proto --go_opt=module=github.com/foxlovecat1989/golang-grpc-prac --go_out=./calculator/proto --go-grpc_opt=module=github.com/foxlovecat1989/golang-grpc-prac --go-grpc_out=./calculator/proto ./calculator/proto/*.proto

calculator-build:
	go build -o bin/calculator/server ./calculator/server
	go build -o bin/calculator/client ./calculator/client

calculator-server-run:
	./bin/calculator/server

calculator-client-run:
	./bin/calculator/client


.PHONY: greet greet-server-run greet-client-run calculator calculator-server-run calculator-client-run greet-proto-gen greet-build calculator-proto-gen calculator-build

