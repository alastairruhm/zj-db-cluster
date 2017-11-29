### Makefile --- 


all: build

build:
	mkdir -p bin
	go build -o bin/zj-db-cluster github.com/alastairruhm/zj-db-cluster

test:
	go test -v ./...

bench:
	go test -bench=. github.com/alastairruhm/influx-proxy/backend

clean:
	rm -rf bin

### Makefile ends here
