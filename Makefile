.DEFAULT_GOAL    := all

EXECUTABLE       := mtzdate
SOURCES          := $(shell echo *.go)
LDFLAGS          := -ldflags="-s -w"
GOMETALINTER_PKG := github.com/alecthomas/gometalinter
VGO_PKG          := github.com/tanakapayam/vgo

GO               := go
GOBIN            ?= $(shell go env GOBIN)
GOFMT            := gofmt
GOGET            := go get
GOMETALINTER     := gometalinter
VGO              := vgo mod
UPX              := upx

DOCKER_PRUNE     := docker system prune --force
DOCKER_RMI       := docker rmi --force tanakapayam/${EXECUTABLE}:latest
DOCKER_BUILD     := docker build --tag tanakapayam/${EXECUTABLE}:latest .
DOCKER_RUN       := docker run --env "MTZDATE_TIMEZONES=$${MTZDATE_TIMEZONES:-UTC}" --env "MTZDATE_FLAGS=$${MTZDATE_FLAGS:-}" --tty tanakapayam/${EXECUTABLE}:latest
DOCKER_RUN_LOOP  := docker run --env "MTZDATE_LOOP=1" --env "MTZDATE_TIMEZONES=$${MTZDATE_TIMEZONES:-UTC}" --env "MTZDATE_FLAGS=$${MTZDATE_FLAGS:-}" --interactive --tty tanakapayam/${EXECUTABLE}:latest

ARCH             := amd64
BUILD            := GOARCH=${ARCH} $(GO) build -i ${LDFLAGS}
FORMAT           := $(GOFMT) -s -w
INSTALL          := $(GO) install
STRIP            := $(UPX) --best --brute

LOCAL            := $(BUILD) -o ${EXECUTABLE}

SHELL            := bash
BOLD             != tput bold
GREEN            != tput setaf 2
ORANGE           != tput setaf 172
RESET            != tput sgr0

.PHONY: all format dep lint $(GOMETALINTER) docker-build docker-run docker-run-loop install

all: ${SOURCES} ${EXECUTABLE}

format:
	$(FORMAT) ${SOURCES}
	@echo

dep:
	$(GOGET) -u ${VGO_PKG}
	$(VGO) -init || true
	$(VGO) -sync
	$(VGO) -vendor
	@echo

lint: $(GOMETALINTER)
	$(GOGET) ${GOMETALINTER_PKG}
	$(GOMETALINTER) --vendor ./... || true
	@echo

$(GOMETALINTER):
	$(GOGET) ${GOMETALINTER_PKG}
	$(GOMETALINTER) --install &> /dev/null
	@echo

${EXECUTABLE}: format dep lint $(GOMETALINTER)
	$(LOCAL)
	@echo

docker-build:
	@echo "${BOLD}${GREEN}# building docker image...${RESET}"
	@echo
	@echo "${BOLD}${ORANGE}## $(DOCKER_PRUNE)${RESET}"
	@$(DOCKER_PRUNE) || true
	@echo
	@echo "${BOLD}${ORANGE}## $(DOCKER_RMI)${RESET}"
	@$(DOCKER_RMI) || true
	@echo
	@echo "${BOLD}${ORANGE}## $(DOCKER_BUILD)${RESET}"
	@$(DOCKER_BUILD)
	@echo
	@echo "${BOLD}${ORANGE}## $(DOCKER_RUN)${RESET}"
	@$(DOCKER_RUN)
	@echo

docker-run:
	@$(DOCKER_RUN)

docker-run-loop:
	@$(DOCKER_RUN_LOOP)

install: all
	$(INSTALL) .
	@echo

install-strip: install
	$(STRIP) ${GOBIN}/${EXECUTABLE}
	@echo
