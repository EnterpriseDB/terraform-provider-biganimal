
TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=hashicorp.com
NAMESPACE=edu
NAME=biganimal
BINARY=terraform-provider-${NAME}
VERSION=0.3.1
OS_ARCH=darwin_amd64

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
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

.PHONY: docs
docs:
	go generate

.PHONY: tf-plan
tf-plan:
	terraform -chdir=examples/provider plan

.PHONY: tf-apply
tf-apply: 
	terraform -chdir=examples/provider apply -auto-approve

