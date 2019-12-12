
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

lint:
	go vet ./... && \
	golangci-lint run --enable-all \
		-D funlen \
		-D gochecknoglobals \
		-D gochecknoinits \
		-D godox \
		-D lll \
		-D nakedret \
		-D wsl \
		./...

.PHONY: gorest test
