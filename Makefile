.PHONEY: all

all: build_all

build_all:
	+$(MAKE) -C ./frontend
	+$(MAKE) -C ./go_ticket_microservice
	+$(MAKE) -C ./go_user_microservice

