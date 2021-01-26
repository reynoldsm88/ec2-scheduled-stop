.PHONY: test

#test:
#	rm -r -f target/test
#	go test ./... -v -short

clean:
	rm -rf target/*

build: # testing
	go build -o target/ec2-scheduled-stop main.go

build-linux:
	GOOS=linux
	go build -o target/ec2-scheduled-stop

package-lambda: build
	zip -r -j target/function.zip target/ec2-scheduled-stop
