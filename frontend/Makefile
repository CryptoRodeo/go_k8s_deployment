all: build_and_publish

build_and_publish: \
	build \
	publish 

build:
	@echo "------------------------------"
	@echo "Building front end"
	@echo "------------------------------"
	npm run build && \
	docker build . -t kilgorevonnegut/microservice_frontend:v1

publish:
	@echo "WIP :)"
