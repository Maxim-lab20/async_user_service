# Show Go installation link
install-go:
	@echo "Install Go from https://go.dev/dl/ manually depending on your OS."

# Show Docker Desktop installation link
install-docker:
	@echo "Install Docker Desktop from https://www.docker.com/products/docker-desktop/ depending on your OS."

# Start Redis and Postgres using Docker Compose command on the file within /setup folder
start-docker:
	docker-compose -f setup/docker-compose.yml up -d

# Build the Go app (works both for Windows and Linux/Mac)
build:
	go build -o app .

# Run the Go app (smart run for both Windows and Unix)
run:
ifeq ($(OS),Windows_NT)
	./app.exe
else
	./app
endif
