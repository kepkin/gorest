
gorest: install
	go build ./cmd/gorest

test: install
	go test -count=1 ./...

install:
	go get -d ./...

generate:
	go generate ./...

fmt:
	go fmt ./...

.PHONY: gorest test
