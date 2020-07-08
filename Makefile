.PHONY: all run-local bench docker coverage.xml

all: go-benchalign

SRC := $(wildcard go.mod go.sum *.go **/*.go)

go-benchngs: $(SRC)
	GOOS=darwin go build -o go-benchngs .

run-local:
	go run main.go

bench:
	go test -v -run=XXX -bench=. ./...

docker:
	go mod vendor

coverage.xml:
	go run github.com/axw/gocov/gocov test $(shell go list ./... ) -coverpkg=$(shell go list ./... | tr '\n' ',') | \
	go run github.com/AlekSi/gocov-xml > coverage.xml

%:
	@:
