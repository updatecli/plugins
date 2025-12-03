build:
	GOOS=wasip1 GOARCH=wasm go build -tags std -o plugin.wasm .

test: build
	go test ./...

