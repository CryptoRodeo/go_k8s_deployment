.PHONEY: all
SHELL=/bin/bash
DIRS = frontend go_ticket_microservice go_user_microservice
MYDIR = $(shell pwd)
all: build_all

list:
	$(foreach dir, $(DIRS), +$(MAKE) --directory $(MYDIR)/$(dir))

build_all:
	+$(MAKE) -C ./frontend
	+$(MAKE) -C ./go_ticket_microservice
	+$(MAKE) -C ./go_user_microservice

publish_all:
	+$(MAKE) publish -C ./frontend
	+$(MAKE) publish -C ./go_ticket_microservice
	+$(MAKE) publish -C ./go_user_microservice
