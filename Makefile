run:
	docker compose up -d

mocks:
	mockgen -source=model/model.go -destination=test/mocks/mocks.go -package=mocks

test:
	go test -count=1 ./...

.PHONY: run mocks test