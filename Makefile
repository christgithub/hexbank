
deps: tools
	go mod download

tools:
	# Install mockgen
	go get github.com/golang/mock/mockgen@v1.4.0

.PHONY: mocks
mocks:
	rm -rf mocks
	go generate ./...