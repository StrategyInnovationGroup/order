.PHONY: build run stop clean

DOCKER_COMPOSE := docker-compose.yaml

build:
	docker-compose -f $(DOCKER_COMPOSE) build
	go build .

run:
	docker-compose -f $(DOCKER_COMPOSE) up -d

gorun:
	bash ./go-run.sh
    
stop:
	docker-compose -f $(DOCKER_COMPOSE) down

clean: stop
	docker-compose -f $(DOCKER_COMPOSE) rm -f