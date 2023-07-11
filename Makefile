build:
	@go build -o bin/qrcrypt -v
	@echo "Build complete"
run:
	@go run .
clean:
	@rm -f bin/qrcrypt
	@echo "Clean complete"
test:
	@go test -v ./...
	@echo "Test complete"
	