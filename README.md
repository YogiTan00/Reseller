# Reseller
Management Sales https://github.com/YogiTan00/Reseller

# Prerequisites
1. golang (v1.22.3 or above) | https://go.dev/doc/install
2. Mysql database
3. Protobuf Compiler | https://github.com/protocolbuffers/protobuf
   - check program is it already installed => `protoc --version`
4. Buf 
   - `go install github.com/bufbuild/buf/cmd/buf@v1.28.1`
   - check program is it already installed => `buf --version`
5. Mockery | https://github.com/vektra/mockery (v2.39.1 or above)
   - `go install github.com/vektra/mockery/v2@v2.39.1`
   - check program is it already installed => `mockery --version`

# Grpc
1. sync go modules -> `go mod tidy`
2. download GRPC ecosystem modules -> \
   `go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway`\
   `go get google.golang.org/protobuf/cmd/protoc-gen-go`\
   `go get google.golang.org/grpc/cmd/protoc-gen-go-grpc`
3. install GRPC ecosystem modules -> \
   `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway`\
   `go install google.golang.org/protobuf/cmd/protoc-gen-go`\
   `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`

# Migration Mysql
1. download and install golang-migrate -> https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
2. migrate up service (e.g., service auth) -> \
   `migrate -database "mysql://root:root@tcp(localhost:3306)/reseller" -path services/migrations/mysql up`
3. create new migration (if needed) -> \
   `migrate create -ext sql -dir services/migrations/mysql create_product`