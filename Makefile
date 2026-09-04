all: run
run:
	@go run .
deps:
	@go mod tidy
check:
	@go vet .
debug:
	@dlv debug --headless --listen=127.0.0.1:43000 --api-version=2
debugc:
	@dlv connect 127.0.0.1:43000
