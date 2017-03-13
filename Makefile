GOSRC = ./../..
EUDATSRC = ./..
WEBUI = frontend/webui
INTERNALSERVICES = services/_internal

build: dependencies
	$(GOPATH)/bin/golint ./...
	go vet ./...
	go test ./...
	go build ./...

dependencies: $(GOSRC)/golang/lint/golint $(GOSRC)/gorilla/mux $(GOSRC)/pborman/uuid

$(GOSRC)/golang/lint/golint:
	go get -u github.com/golang/lint/golint

$(GOSRC)/gorilla/mux:
	go get -u github.com/gorilla/mux

$(GOSRC)/pborman/uuid:
	go get -u github.com/pborman/uuid

run:
	go run main.go

clean:
	go clean ./...

.PHONY: build dependencies run clean
