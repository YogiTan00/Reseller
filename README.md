# Reseller
Management Sales https://github.com/YogiTan00/Reseller

### Prerequisites
1. Golang (v1.21 or Above)
   - https://go.dev/doc/install
2. Mysql Database
3. Protobuf Compiler
   - https://github.com/protocolbuffers/protobuf | https://github.com/bufbuild/buf/releases
   - check program is it already installed => `protoc --version`
4. Buf
   - https://connectrpc.com/docs/go/getting-started/ (v1.34.0 or above)
   - `go install github.com/bufbuild/buf/cmd/buf@latest`
   - check program is it already installed => `buf --version`
5. Mockery
   - https://github.com/vektra/mockery (v2.39.1 or above)
   - `go install github.com/vektra/mockery/v2@latest`
   - check program is it already installed => `mockery --version`

### Grpc
1. Sync Go Modules -> `go mod tidy`
2. Download GRPC Ecosystem Nodules -> \
   `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`\
   `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`\
   `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest` \
   Check \
   `where protoc-gen-go`\
   `where protoc-gen-go-grpc`\
   `where protoc-gen-grpc-gateway`
3. Install GRPC Ecosystem Modules -> \
   `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway`\
   `go install google.golang.org/protobuf/cmd/protoc-gen-go`\
   `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`

### Migration Mysql
1. Download and Install Golang-Migrate -> https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
2. Migrate Up Service (Optional) -> \
   `migrate -database "mysql://root:root@tcp(localhost:3306)/reseller" -path services/migrations/mysql up`
3. Create New Migration (Optional) -> \
   `migrate create -ext sql -dir services/migrations/mysql create_product`

### Commit Message
1. Feat: Add Feature
2. Fix: Fixing Fiture
3. Test: Make or Increased Unitest