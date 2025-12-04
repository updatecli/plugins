build:
	GOOS=wasip1 GOARCH=wasm go build -tags std -o demo.wasm .

test: build
	go test ./...

