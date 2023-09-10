# ##############################################################################
# # File: Makefile                                                             #
# # Project: epub2pdf                                                          #
# # Created Date: 2023/09/10 16:56:37                                          #
# # Author: realjf                                                             #
# # -----                                                                      #
# # Last Modified: 2023/09/11 01:33:31                                         #
# # Modified By: realjf                                                        #
# # -----                                                                      #
# # Copyright (c) 2023 realjf                                                  #
# ##############################################################################


# ========================================== project ==========================================

VERSION="0.1.0"
APP_NAME="epub2pdf"
ARCH=amd64
# target
GOOS=linux



BIN=bin/${APP_NAME}-linux-${ARCH}-${VERSION}
BIN_WIN=bin/${APP_NAME}-windows-${ARCH}-${VERSION}.exe
BIN_MACOS=bin/${APP_NAME}-darwin-${ARCH}-${VERSION}



.PHONY: build
build:
	@if test "$(GOOS)" = "linux"; then \
    	env CGO_ENABLED=1 GOOS=linux GOARCH=${ARCH} go build -ldflags '-s -w -X main.Version=${VERSION}' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN} ./app/main.go; \
	elif test "${GOOS}" = "windows"; then \
    	env CGO_ENABLED=1 GOOS=windows GOARCH=${ARCH} go build -ldflags '-s -w -X main.Version=${VERSION} -H=windowsgui -extldflags "-static"' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN_WIN} ./app/main.go; \
	elif test "${GOOS}" = "darwin"; then \
    	env CGO_ENABLED=1 GOOS=darwin GOARCH=${ARCH} go build -ldflags '-s -w -X main.Version=${VERSION}' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN_MACOS} ./app/main.go; \
	else \
    	echo "Target OS is not supported"; \
	fi
	@echo 'done';



.PHONY: lint
lint:
	@golangci-lint run -v ./...




# ========================================== fyne ==========================================

# install fyne tool
.PHONY: install
install:
	@go get fyne.io/fyne/v2@latest
	@go install fyne.io/fyne/v2/cmd/fyne@latest

#
.PHONY: install_cross
install_cross:
	@go install github.com/fyne-io/fyne-cross@latest


.PHONY: deps
deps:
	@sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev -y
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.1







# ========================================== git ==========================================


B ?= master

.PHONY: push
push:
	@git add -A && git commit -m "update" && git push origin ${B}


# make tag t=<your_version>
.PHONY: tag
tag:
	@echo '${t}'
	@git tag -a ${t} -m "${t}" && git push origin ${t}



.PHONY: dtag
dtag:
	@echo 'delete ${t}'
	@git push --delete origin ${t} && git tag -d ${t}


