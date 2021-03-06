#Copyright Uhuchain. All Rights Reserved.
#
#SPDX-License-Identifier: Apache-2.0
#
#
# Supported Targets:
# all : runs unit and integration tests
# depend: checks that test dependencies are installed
# depend-install: installs test dependencies
# unit-test: runs all the unit tests
# integration-test: runs all the integration tests
# checks: runs all check conditions (license, spelling, linting)
# clean: stops docker conatainers used for integration testing
# mock-gen: generate mocks needed for testing (using mockgen)
# channel-config-gen: generates the channel configuration transactions and blocks used by tests
# populate: populates generated files (not included in git) - currently only vendor
# populate-vendor: populate the vendor directory based on the lock
# populate-clean: cleans up populated files (might become part of clean eventually) 
# thirdparty-pin: pulls (and patches) pinned dependencies into the project under internal
#

# Tool commands (overridable)
GO_CMD             ?= go
GO_DEP_CMD         ?= dep
DOCKER_CMD         ?= docker
DOCKER_COMPOSE_CMD ?= docker-compose

# Build flags (overridable)
GO_LDFLAGS                 ?= -ldflags=-s
GO_TESTFLAGS               ?=
FABRIC_SDK_EXPERIMENTAL    ?= true
FABRIC_SDK_EXTRA_GO_TAGS   ?=
FABRIC_SDK_POPULATE_VENDOR ?= true

# Fabric tool versions (overridable)
FABRIC_TOOLS_VERSION ?= 1.0.3
FABRIC_BASE_VERSION  ?= 0.4.2

# Fabric base docker image (overridable)
FABRIC_BASE_IMAGE   ?= hyperledger/fabric-baseimage
FABRIC_BASE_TAG     ?= $(ARCH)-$(FABRIC_BASE_VERSION)

# Fabric tools docker image (overridable)
FABRIC_TOOLS_IMAGE  ?= hyperledger/fabric-tools
FABRIC_TOOLS_TAG    ?= $(ARCH)-$(FABRIC_TOOLS_VERSION)

# Upstream fabric patching (overridable)
THIRDPARTY_FABRIC_CA_BRANCH ?= master
THIRDPARTY_FABRIC_CA_COMMIT ?= 2886abda6792cf3b5e708ed18dbde07106597071
THIRDPARTY_FABRIC_BRANCH    ?= master
THIRDPARTY_FABRIC_COMMIT    ?= f754f40d3165571cecf5fce43c8a034559983311

# Local variables used by makefile
PACKAGE_NAME := github.com/uhuchain/uhuchain
ARCH         := $(shell uname -m)

# The version of dep that will be installed by depend-install (or in the CI)
GO_DEP_COMMIT := v0.3.1

# Setup Go Tags
GO_TAGS := $(FABRIC_SDK_EXTRA_GO_TAGS)
ifeq ($(FABRIC_SDK_EXPERIMENTAL),true)
GO_TAGS += experimental
endif

# Detect CI
ifdef JENKINS_URL
export FABRIC_SDKGO_DEPEND_INSTALL := true
endif

# Global environment exported for scripts
export GO_CMD
export GO_DEP_CMD
export ARCH
export GO_LDFLAGS
export GO_DEP_COMMIT
export GO_TAGS
export GO_TESTFLAGS

all: checks unit-test integration-test

depend:
	@test/scripts/dependencies.sh

depend-install:
	@FABRIC_SDKGO_DEPEND_INSTALL="true" test/scripts/dependencies.sh

checks: depend license lint spelling

.PHONY: license build-softhsm2-image
license:
	@test/scripts/check_license.sh

lint:
	@test/scripts/check_lint.sh

spelling:
	@test/scripts/check_spelling.sh

unit-test: checks
	$(GO_CMD) test -v ./test/unit

unit-tests: unit-test

prepare-network:
	@echo "=========== Preparing network ==========="
	@cd ./test/uhuchain-network-dev/scripts && chmod +x prepare.sh && ./prepare.sh car-ledger 1 prepare 1.0

exec-integration-test: prepare-network
	@echo "=========== Executing integration tests ==========="
	@cd ./test/integration && $(GO_CMD) test -v

integration-test: clean depend populate
	@echo "=========== Starting uhuchain test network ==========="
	@cd ./test/uhuchain-network-dev && $(DOCKER_COMPOSE_CMD) -f docker-compose.yaml up -d 
	@docker exec -it uhuchain-server bash -c 'make test-and-run-server'
	@docker logs -f uhuchain-server

channel-config-gen:
	@echo "Generating test channel configuration transactions and blocks ..."
	@$(DOCKER_CMD) run -i \
		-v $(abspath .):/opt/gopath/src/$(PACKAGE_NAME) \
		$(FABRIC_TOOLS_IMAGE):$(FABRIC_TOOLS_TAG) \
		/bin/bash -c "/opt/gopath/src/${PACKAGE_NAME}/test/scripts/generate_channeltx.sh"

thirdparty-pin:
	@echo "Pinning third party packages ..."
	@UPSTREAM_COMMIT=$(THIRDPARTY_FABRIC_COMMIT) UPSTREAM_BRANCH=$(THIRDPARTY_FABRIC_BRANCH) scripts/third_party_pins/fabric/apply_upstream.sh
	@UPSTREAM_COMMIT=$(THIRDPARTY_FABRIC_CA_COMMIT) UPSTREAM_BRANCH=$(THIRDPARTY_FABRIC_CA_BRANCH) scripts/third_party_pins/fabric-ca/apply_upstream.sh

populate: populate-vendor

populate-vendor:
ifeq ($(FABRIC_SDK_POPULATE_VENDOR),true)
	@echo "Populating vendor ..."
	@$(GO_DEP_CMD) ensure -vendor-only
endif

populate-clean:
	rm -Rf vendor

build-linux:
	@echo "Building uhuchain-server executable for linux ..."
	env GOOS=linux GOARCH=amd64 $(GO_CMD) build -o build/linux/uhuchain-server ./cmd/uhuchain-server

install-server:
	@echo "=========== Install Uhuchain server ==========="
	@cd ./cmd/uhuchain-server && $(GO_CMD) install

test-and-run-server: exec-integration-test install-server
	@echo "=========== Execute Uhuchain server ==========="
	uhuchain-server --scheme=http --host=0.0.0.0 --port=3333 

clean:
	$(GO_CMD) clean
	rm -Rf /tmp/enroll_user /tmp/msp /tmp/keyvaluestore /tmp/hfc-kvs
	rm -f integration-report.xml report.xml
	@cd ./test/uhuchain-network-dev && $(DOCKER_COMPOSE_CMD) -f docker-compose.yaml down
	./test/scripts/clean_docker_images.sh uhuchain
