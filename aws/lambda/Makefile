.PHONY: build clean deploy

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/hello ./hello

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
