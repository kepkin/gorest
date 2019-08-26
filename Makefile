
gorest: generate
	go build ./cmd/gorest

generate:
	go generate ./...

fmt:
	go fmt ./...

test:
	go test ./...

.PHONY: gorest test
