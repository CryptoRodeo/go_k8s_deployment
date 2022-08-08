.PHONY: all

image_tag := kilgorevonnegut/go-ticket-microservice:v1
name := "Go Ticket Microservice"

all: build

build:
	@echo "------------------------------"
	@echo "Building ${name}"
	@echo "------------------------------"
	docker build . -t ${image_tag}

publish:
	@echo "------------------------------"
	@echo "Publishing ${name} to Docker"
	@echo "------------------------------"
	docker push ${image_tag}
