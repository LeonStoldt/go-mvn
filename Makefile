PROJECT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
SHELL = /bin/bash

run:
	@cd ${PROJECT_DIR}
	go run main.go

test:
	@cd ${PROJECT_DIR}
	# best test we have so far
	go run main.go
