run/docker:
	docker compose up -d

run/local:
	go build ./cmd/mjurl/main.go && ./main

mocks:
	mockgen -source=model/model.go -destination=test/mocks/mocks.go -package=mocks

test:
	go test -count=1 ./...

clean:
	rm main

grpc:
	protoc --go_out=./internal/transport/grpc/pb --go_opt=paths=source_relative \
    	--go-grpc_out=./internal/transport/grpc/pb --go-grpc_opt=paths=source_relative \
    	mjurl.proto

.PHONY: run mocks test