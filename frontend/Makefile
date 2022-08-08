.PHONY: all

image_tag := kilgorevonnegut/microservice_frontend:v1
name := "TicketMaster frontend"

all: build

build:
	@echo "------------------------------"
	@echo "Building ${name}"
	@echo "------------------------------"
	npm install && \
	npm run build && \
	docker build . -t ${image_tag}

publish:
	@echo "------------------------------"
	@echo "Publishing ${name} to Docker"
	@echo "------------------------------"
	docker push ${image_tag}
