# Project Name

SHA1         		:= $(shell git rev-parse --verify --short HEAD)
MAJOR_VERSION		:= $(shell cat lib.json | sed -n 's/.*"major": "\(.*\)",/\1/p')
MINOR_VERSION		:= $(shell cat lib.json | sed -n 's/.*"minor": "\(.*\)"/\1/p')
INTERNAL_BUILD_ID	:= $(shell [ -z "${TRAVIS_BUILD_NUMBER}" ] && echo "local" || echo ${TRAVIS_BUILD_NUMBER})
BINARY				:= $(shell cat service.json | sed -n 's/.*"name": "\(.*\)",/\1/p')
VERSION				:= $(shell echo "${MAJOR_VERSION}_${MINOR_VERSION}_${INTERNAL_BUILD_ID}_${SHA1}")
BUILD_IMAGE			:= $(shell echo "golang:1.11.4")
PWD					:= $(shell pwd)

.DEFAULT_GOAL := test

.PHONY: test
test:
	@echo "Running Tests"

	docker run -e GO111MODULE=auto --rm -t -v $(PWD):/usr/src/myapp -w /usr/src/myapp $(BUILD_IMAGE) sh -c "go test -cover -v ./... -count=1"
	@echo "Completed tests"


