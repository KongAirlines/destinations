include base.mk

destinations_mkfile_dir := $(CURRDIR)

check-dependencies:
	@$(call check-dependency,go)
	@$(call check-dependency,jq)
	@$(call check-dependency,deck)
	@$(call check-dependency,docker)

test: check-dependencies
	@go test -v ./...

build: check-dependencies
	@go generate ./...
	@go build .

build-docker:
	@docker build -t kongair/destinations:latest .

run: check-dependencies build
	@./destinations ${KONG_AIR_DESTINATIONS_PORT}

docker: build-docker
	@docker run -d --name kongair-destinations -p ${KONG_AIR_DESTINATIONS_PORT}:${KONG_AIR_DESTINATIONS_PORT} kongair/destinations:latest

kill-docker:
	-@docker stop kong-air-destinations-svc
	-@docker rm kong-air-destinations-svc
	@if [ $$? -ne 0 ]; then $(call echo_fail,Failed to kill the docker containers); exit 1; else $(call echo_pass,Killed the docker container); fi

