
gorest: generate
	go build ./cmd/gorest

generate:
	go generate ./...

fmt:
	go fmt ./...

test: generate
	go test ./... -count=1

.PHONY: gorest test
