GOTOOLS = \
	github.com/mitchellh/gox \
	github.com/Masterminds/glide
GOTOOLS_CHECK = gox glide
PACKAGES=$(shell go list ./... | grep -v '/vendor/')
BUILD_TAGS?=trading
TMHOME = $${TMHOME:-$$HOME/.trading}
# BUILD_FLAGS = -ldflags "-X github.com/epappas/trading/version.GitCommit=`git rev-parse --short HEAD`"
# BUILD_FLAGS = ""

all: build build_race install

########################################
### Build

build:
	go build $(BUILD_FLAGS) -o build/trading ./cmd/trading/

build_race:
	go build -race $(BUILD_FLAGS) -o build/trading ./cmd/trading

install:
	go install $(BUILD_FLAGS) ./cmd/trading


########################################
### Testing

test:
	@echo "--> Running go test"
	@go test $(PACKAGES)

test_race:
	@echo "--> Running go test --race"
	@go test -v -race $(PACKAGES)

test_integrations:
	@bash ./test/test.sh

test_release:
	@go test -tags release $(PACKAGES)

test100:
	@for i in {1..100}; do make test; done


# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: build build_race install
