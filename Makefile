run:
	go run ./cmd/mjurl/main.go

mocks:
	mockgen -source=internal/url/service.go -destination=test/mocks/mocks.go -package=mocks