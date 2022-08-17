.PHONY: all

image_tag := kilgorevonnegut/go-user-microservice:v1
name := "Go User Microservice"

all: build

build:
	@echo "------------------------------"
	@echo "Building ${name}"
	@echo "------------------------------"
	docker build . -t ${image_tag} --network=host

publish:
	@echo "------------------------------"
	@echo "Publishing ${name} to Docker"
	@echo "------------------------------"
	docker push ${image_tag}
