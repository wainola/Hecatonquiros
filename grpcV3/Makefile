compile:
	protoc -I proto/ proto/list.proto --go_out=plugins=grpc:proto
cqlsh:
	docker exec -it grpc_demo  bash -c "cqlsh"
docker:
	docker-compose up -d
docker-down:
	docker-compose down
tables:
	go run ./cmd/server/db/main.go