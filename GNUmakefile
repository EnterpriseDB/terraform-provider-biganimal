
TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=registry.terraform.io
NAMESPACE=EnterpriseDB
NAME=biganimal
BINARY=terraform-provider-${NAME}
VERSION=0.7.0

# Figure out the OS and ARCH of the
# builder machine
os       ?= $(shell uname|tr A-Z a-z)
ifeq ($(shell uname -m),x86_64)
  arch   ?= amd64
endif
ifeq ($(shell uname -m),i686)
  arch   ?= 386
endif
ifeq ($(shell uname -m),aarch64)
  arch   ?= arm
endif
ifeq ($(shell uname -m),arm64)
  arch   ?= arm64
endif
OS_ARCH=${os}_${arch}

default: install

# Run acceptance tests
.PHONY: testacc


build:
	go build -o ${BINARY}

release:
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 godotenv go test $(TEST) -v $(TESTARGS) -run ".*?$(TYPE)$(NAME).*?" -timeout 4h

.PHONY: docs
unexport BA_BEARER_TOKEN
unexport BA_API_URI
docs:
	go generate

.PHONY: tf-plan
tf-plan:
	terraform -chdir=examples/provider plan

.PHONY: tf-apply
tf-apply:
	terraform -chdir=examples/provider apply -auto-approve
