.PHONEY: all
all: build_all

build_all:
	+$(MAKE) -C ./frontend
	+$(MAKE) -C ./go_ticket_microservice
	+$(MAKE) -C ./go_user_microservice

publish_all:
	+$(MAKE) publish -C ./frontend
	+$(MAKE) publish -C ./go_ticket_microservice
	+$(MAKE) publish -C ./go_user_microservice
