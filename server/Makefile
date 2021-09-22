clean:
	rm -rf build

build: clean
	go build -o build/tulkas

run:
	go run .

test:
	go test ./... -v

lint:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.38.0 golangci-lint run -v

fmt:
	docker run --rm -v $(PWD):/data cytopia/gofmt -s -w .

cover:
	BOMBUR_DB_URI="postgresql://localhost/bombur?user=bombur&password=bombur" go test cover ./...