### Makefile --- 


all: build

build:
	mkdir -p bin
	go build -o bin/zj-db-cluster github.com/alastairruhm/zj-db-cluster

lint:
	gometalinter --config=gometalinter_config.json ./...
	vendorcheck ./...

test:
	go test -v ./...

clean:
	rm -rf bin

integration-test:
	go test -tags=integration -v

### Makefile ends here
