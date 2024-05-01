build:
	go build -o bin/ccwc ./cmd/...

run: build
	bin/ccwc -m test.txt
	cat test.txt | bin/ccwc -l
	ls | bin/ccwc

test: 
	go test ./...

usage: build
	bin/ccwc -h
	
