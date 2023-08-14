DOCKER_IMG_NAME := toyrobot-app
DOCKER_CONTAINER_NAME := toyrobot-container
DOCKERFILE_PATH := ./build/Dockerfile

build-docker:
	@docker build -f $(DOCKERFILE_PATH) -t  $(DOCKER_IMG_NAME) .

run-docker: build-docker
	@docker run -it --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMG_NAME) play

test:
	@go mod download
	@go mod verify
	@go test ./... -v

build:
	@go mod download
	@go mod verify
	@go build -o $(DOCKER_IMG_NAME) .

run:
	@go mod download
	@go mod verify
	@go build -o $(DOCKER_IMG_NAME) .
	@./$(DOCKER_IMG_NAME) play

help:
	@echo "Available commands:"
	@echo "- build-docker: Build the Docker image"
	@echo "- run-docker: Run the app inside a Docker container"
	@echo "- test: Run tests"
	@echo "- run: Run the app"

.PHONY: build-docker run-docker build test run help
