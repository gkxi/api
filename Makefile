GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
BRANCH=$(shell git symbolic-ref -q --short HEAD)
REVISION=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date +%FT%T%Z)
PROTO_FILES=$(shell find . -name *.proto)
KRATOS_VERSION=$(shell go mod graph |grep go-kratos/kratos/v2 |head -n 1 |awk -F '@' '{print $$2}')
KRATOS=$(GOPATH)/pkg/mod/github.com/go-kratos/kratos/v2@$(KRATOS_VERSION)
PKG_VERSION=$(shell go mod graph |grep go-bamboo/pkg |head -n 1 |awk -F '@' '{print $$2}')
PKG=$(GOPATH)/pkg/mod/github.com/go-bamboo/pkg@$(PKG_VERSION)
CONTRIB_VERSION=$(shell go mod graph |grep go-bamboo/contrib |head -n 1 |awk -F '@' '{print $$2}')
CONTRIB=$(GOPATH)/pkg/mod/github.com/go-bamboo/contrib@$(CONTRIB_VERSION)
PWD := $(shell pwd)

.PHONY: errors
errors:
	protoc --proto_path=. \
		   --proto_path=$(PKG) \
		   --proto_path=$(KRATOS) \
           --proto_path=$(KRATOS)/api \
           --proto_path=$(KRATOS)/third_party \
           --proto_path=$(PWD)/../third_party \
           --go_out=paths=source_relative:. \
           --go-errors_out=paths=source_relative:. \
           ./address/ecode.proto \
           ./tran/ecode.proto

.PHONY: api
api:
	protoc --proto_path=. \
		   --proto_path=$(PKG) \
		   --proto_path=$(KRATOS) \
           --proto_path=$(KRATOS)/api \
           --proto_path=$(KRATOS)/third_party \
           --proto_path=$(PWD)/../third_party \
           --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
           --go-http_out=paths=source_relative:. \
           --go-errors_out=paths=source_relative:. \
           --validate_out=lang=go,paths=source_relative:. \
           --openapi_out=. \
           ./address/v1/api.proto \
           ./tran/v1/api.proto \
           ./base/types.proto

