vendor:
	@echo "Installing dependencies"
	go get ./...

build: vendor fmt
	@echo "Building the application"
	go build -o expenses

clean:
	@echo "Removing app binary"
	rm -f expenses

fmt:
	@echo "Formatting the code"
	go fmt ./...
