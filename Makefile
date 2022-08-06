all: build_all

build_all: \
	build_frontend \
	build_go_ticket\
	build_go_user

build_frontend:
	cd ./frontend && \
	make	

build_go_ticket:
	cd ./go_ticket_microservice && \
		make	

build_go_user:
	cd ./go_user_microservice && \
		make
