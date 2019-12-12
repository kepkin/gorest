
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
    	-D gochecknoglobals \
    	-D lll \
    	-D gochecknoinits \
    	-D nakedret \
    	./...

.PHONY: gorest test
