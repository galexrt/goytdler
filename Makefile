PROJECTNAME ?= goytdler
DESCRIPTION ?= goytdler - Simple webinterface to use youtube-dl.
MAINTAINER  ?= Alexander Trost <galexrt@googlemail.com>
HOMEPAGE    ?= https://github.com/galexrt/goytdler

GO111MODULE     ?= on
GO              ?= go
FPM             ?= fpm
PROMU           := $(GOPATH)/bin/promu
GOASSETSBUILDER := $(GOPATH)/bin/go-assets-builder
PREFIX          ?= $(shell pwd)
BIN_DIR         ?= $(PREFIX)/.build
TARBALL_DIR     ?= $(PREFIX)/.tarball
PACKAGE_DIR     ?= $(PREFIX)/.package
ARCH            ?= amd64
PACKAGE_ARCH    ?= linux-amd64

# The GOHOSTARM and PROMU parts have been taken from the prometheus/promu repository
# which is licensed under Apache License 2.0 Copyright 2018 The Prometheus Authors
FIRST_GOPATH := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))

GOHOSTOS     ?= $(shell $(GO) env GOHOSTOS)
GOHOSTARCH   ?= $(shell $(GO) env GOHOSTARCH)

ifeq (arm, $(GOHOSTARCH))
	GOHOSTARM ?= $(shell GOARM= $(GO) env GOARM)
	GO_BUILD_PLATFORM ?= $(GOHOSTOS)-$(GOHOSTARCH)v$(GOHOSTARM)
else
	GO_BUILD_PLATFORM ?= $(GOHOSTOS)-$(GOHOSTARCH)
endif

PROMU_VERSION ?= 0.7.0
PROMU_URL     := https://github.com/prometheus/promu/releases/download/v$(PROMU_VERSION)/promu-$(PROMU_VERSION).$(GO_BUILD_PLATFORM).tar.gz
# END copied code

pkgs = $(shell go list ./... | grep -v /vendor/ | grep -v /test/)

DOCKER_IMAGE_NAME ?= goytdler
DOCKER_IMAGE_TAG  ?= $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))

all: format style vet test build

build: promu bindata
	@$(PROMU) build --prefix $(PREFIX)

check_license:
	@OUTPUT="$$(promu check licenses)"; \
	if [[ $$OUTPUT ]]; then \
		echo "Found go files without license header:"; \
		echo "$$OUTPUT"; \
		exit 1; \
	else \
		echo "All files with license header"; \
	fi

docker:
	@echo ">> building docker image"
	@docker build -t "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" .

format:
	go fmt $(pkgs)

package-%: build
	mkdir -p -m0755 $(PACKAGE_DIR)/lib/systemd/system $(PACKAGE_DIR)/usr/bin
	mkdir -p $(PACKAGE_DIR)/etc/sysconfig
	cp .build/goytdler $(PACKAGE_DIR)/usr/bin
	cp systemd/goytdler.service $(PACKAGE_DIR)/lib/systemd/system
	cp systemd/sysconfig.goytdler $(PACKAGE_DIR)/etc/sysconfig/goytdler
	cd $(PACKAGE_DIR) && $(FPM) -s dir -t $(patsubst package-%, %, $@) \
	--deb-user root --deb-group root \
	--name $(PROJECTNAME) \
	--version $(shell cat VERSION) \
	--architecture $(PACKAGE_ARCH) \
	--description "$(DESCRIPTION)" \
	--maintainer "$(MAINTAINER)" \
	--url $(HOMEPAGE) \
	usr/ etc/

promu:
	@echo ">> fetching promu"
	@GOOS="$(shell uname -s | tr A-Z a-z)" \
	GOARCH="$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m)))" \
	$(GO) get -u github.com/prometheus/promu

style:
	@echo ">> checking code style"
	@! gofmt -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

tarball: promu
	@echo ">> building release tarball"
	@$(PROMU) tarball --prefix $(PREFIX) $(BIN_DIR)

test:
	@$(GO) test $(pkgs)

test-short:
	@echo ">> running short tests"
	@$(GO) test -short $(pkgs)

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

go-assets-builder:
	@echo ">> fetching go-assets-builder"
	@GOOS="$(shell uname -s | tr A-Z a-z)" \
	GOARCH="$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m)))" \
	$(GO) get -u github.com/jessevdk/go-assets-builder

bindata: go-assets-builder
	$(GOASSETSBUILDER) -o data/assets.go -p data templates

.PHONY: all build crossbuild docker format package promu style tarball test vet
