run:
	docker compose up -d

mocks:
	mockgen -source=model/model.go -destination=test/mocks/mocks.go -package=mocks

test:
	go test -count=1 ./...

grpc:
	protoc --go_out=./internal/transport/grpc/pb --go_opt=paths=source_relative \
    	--go-grpc_out=./internal/transport/grpc/pb --go-grpc_opt=paths=source_relative \
    	mjurl.proto

.PHONY: run mocks test