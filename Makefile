PROTO_DIR=pkg/proto
OUT_DIR=.
PROTO_FILES=$(wildcard $(PROTO_DIR)/*.proto)

GOPATH_BIN=$(shell go env GOPATH)/bin
PATH_WITH_GOBIN=$(GOPATH_BIN):$(PATH)

all: generate

.PHONY: generate
generate:
	@echo "Gerando código Go a partir dos arquivos .proto..."
	@PATH="$(PATH_WITH_GOBIN)" protoc \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

.PHONY: install-tools
install-tools:
	@echo "Instalando plugins protoc-gen-go, protoc-gen-go-grpc e grpcurl... "
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

.PHONY: build
build:
	@echo "Compilando aplicação Go..."
	@go build -o main cmd/main.go
	
.PHONY: clean
clean:
	@echo "Limpando arquivos gerados..."
	@find $(PROTO_DIR) -name '*.pb.go' -delete
