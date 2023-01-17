.PHONEY: all
all: build_all publish_all

.PHONY: build_all
build_all:
	$(MAKE) --directory ./frontend
	$(MAKE) --directory ./go_ticket_microservice
	$(MAKE) --directory ./go_user_microservice

.PHONEY: publish_all
publish_all:
	$(MAKE) publish --directory ./frontend
	$(MAKE) publish --directory ./go_ticket_microservice
	$(MAKE) publish --directory ./go_user_microservice
