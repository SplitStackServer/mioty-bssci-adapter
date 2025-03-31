.PHONY: requirements msg cmd rsp

PROTOC_ARGS := -I=. --go_out=../internal/api --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative

all: requirements msg cmd rsp

requirements:
	go mod download
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

msg:
	protoc ${PROTOC_ARGS} msg/endnode.proto
	protoc ${PROTOC_ARGS} msg/basestation.proto

cmd:
	protoc ${PROTOC_ARGS} cmd/cmd.proto

rsp:
	protoc ${PROTOC_ARGS} rsp/rsp.proto

