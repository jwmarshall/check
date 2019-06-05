SHELL = /usr/bin/env bash
NAME = check
LDFLAGS = -s -w -v=2
STACK_FILE = develop-stack.yml
VERSION = $(shell git describe --always --long)

export NAME
export LDFLAGS
export STACK_FILE
export VERSION

deps:
	go get .

run:
	go run main.go

build: deps
	go build

build-prod: deps
	go build -ldflags="$$LDFLAGS -X check/cmd/version.Version=$$VERSION"

test:
	go test ./...

swarm-init:
	docker swarm init || true

test-stack:
	docker stack deploy -c "$$STACK_FILE" "$$NAME"-testing

stack-pull:
	for i in $$(grep image "$$STACK_FILE" | awk -F':' '{ print $2 }'); do \
		docker pull $$i; \
	done

stack-up:
	docker stack deploy -c "$$STACK_FILE" "$$NAME"-testing

stack-down:
	docker stack rm "$$NAME"-develop