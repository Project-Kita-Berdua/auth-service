proto: 
	protoc pkg/pb/*.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false
server: 
	go run cmd/main.go
postgres: 
	docker run --name grpc-auth-svc -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine