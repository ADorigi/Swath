run:
	@SWATH_VERSION=0.0.0 ./swath

build:
	@go build -v ./...

test:
	@go test -v ./...

# build and run
brun:
	@go build -v ./...
	@SWATH_VERSION=0.0.0 ./swath