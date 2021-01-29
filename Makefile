.PHONY: test

#test:
#	rm -r -f target/test
#	go test ./... -v -short

clean:
	rm -rf target/*

build: # testing
	go build -o target/main main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o target/main

package-lambda: build-linux
	zip -r -j target/function.zip target/main
