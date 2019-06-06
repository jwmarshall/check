SHELL = /usr/bin/env bash
PACKAGE = github.com/jwmarshall
NAME = check
LDFLAGS = -s -w -v=2
STACK_FILE = develop-stack.yml
VERSION = $(shell git describe --always --long)

export PACKAGE
export NAME
export LDFLAGS
export STACK_FILE
export VERSION

deps:
	go get ./...

run:
	go run main.go

build: deps
	go build

build-prod: deps
	go build -ldflags="$$LDFLAGS -X $$PACKAGE/$$NAME/cmd/version.version=$$VERSION"

test:
	go test ./...

docker-image:
	docker build -t "$$NAME" .

swarm-init:
	docker swarm init || true

stack-pull:
	for i in $$(grep image "$$STACK_FILE" | awk -F':' '{ print $2 }'); do \
		docker pull $$i; \
	done

stack-up:
	docker stack deploy -c "$$STACK_FILE" "$$NAME"-develop

stack-down:
	docker stack rm "$$NAME"-develop