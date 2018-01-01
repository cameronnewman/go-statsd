# Project Name

SHA1         		:= $(shell git rev-parse --verify --short HEAD)
MAJOR_VERSION		:= $(shell cat lib.json | sed -n 's/.*"major": "\(.*\)",/\1/p')
MINOR_VERSION		:= $(shell cat lib.json | sed -n 's/.*"minor": "\(.*\)"/\1/p')
INTERNAL_BUILD_ID	:= $(shell [ -z "${BUILD_ID}" ] && echo "local" || echo ${BUILD_ID})
BINARY				:= $(shell cat lib.json | sed -n 's/.*"name": "\(.*\)",/\1/p')
VERSION				:= $(shell echo "${MAJOR_VERSION}_${MINOR_VERSION}_${INTERNAL_BUILD_ID}_${SHA1}")
BUILD_IMAGE			:= $(shell echo "lib-build")
PWD					:= $(shell pwd)

DOT:= .
DASH:= -
# replace . with -
PROJECT= $(subst $(DOT),$(DASH),$(BINARY))

.PHONY: setup
setup:

	@echo $(VERSION)
	@echo $(BINARY)

	docker rmi -f $(BUILD_IMAGE)
	docker build -t=$(BUILD_IMAGE) build/container/

.PHONY: tests
tests: setup

	docker run --rm --name=$(BUILD_IMAGE) -t -v $(PWD):/go/src/github.com/cameronnewman/$(BINARY) -w /go/src/github.com/cameronnewman/$(BINARY) $(BUILD_IMAGE) go test -v ./...
	@echo "Completed tests"
