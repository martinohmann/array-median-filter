default: run

.PHONY: deps test bench run build

deps:
	glide install

test:
	go test $$(go list ./... | grep -v /vendor/)

bench:
	cd pkg/finder && go test -bench=.

run:
	go run main.go

build:
	go build
