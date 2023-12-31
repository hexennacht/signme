GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find * -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find ../../contracts/proto -name *.proto)
endif

.PHONY: api
# generate api proto
api:
	protoc --proto_path=../../contracts \
	       --proto_path=../../contracts/third_party \
 	       --go_out=paths=source_relative:./grpc/auth \
 	       --go-http_out=paths=source_relative:./grpc/auth \
 	       --go-grpc_out=paths=source_relative:./grpc/auth \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
			../../contracts/proto/auth.proto \

	protoc --proto_path=../../contracts \
		   --proto_path=../../contracts/third_party \
		   --go_out=paths=source_relative:./grpc/user \
		   --go-http_out=paths=source_relative:./grpc/user \
		   --go-grpc_out=paths=source_relative:./grpc/user \
		   --openapi_out=fq_schema_naming=true,default_response=false:. \
			../../contracts/proto/user.proto \

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: all
# generate all
all:
	make api;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
