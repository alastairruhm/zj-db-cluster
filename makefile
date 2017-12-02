### Makefile --- 
	
# https://stackoverflow.com/questions/3931741/why-does-make-think-the-target-is-up-to-date
.PHONY: all test clean


all: build

build:
	mkdir -p bin
	go build -o bin/zj-db-cluster github.com/alastairruhm/zj-db-cluster

lint:
	gometalinter --config=gometalinter_config.json ./...
	vendorcheck ./...

test:
	go test -v -gcflags=-l ./...

clean:
	rm -rf bin


### Makefile ends here
