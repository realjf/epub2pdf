VERSION="0.0.2"
BIN="bin/epub2pdf"
BIN_WIN=".exe"
BIN_MACOS=""
ARCH="amd64"

build:
	@env CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build -ldflags '-s -w -X main.Version=${VERSION}' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN}-linux-${ARCH}-${VERSION} main.go
	@echo 'done'

build_win:
	@env CGO_ENABLED=0 GOOS=windows GOARCH=${ARCH} go build -ldflags '-s -w -X main.Version=${VERSION}' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN}-windows-${ARCH}-${VERSION}${BIN_WIN} main.go
	@echo 'done'

build_darwin:
	@env CGO_ENABLED=0 GOOS=darwin GOARCH=${ARCH} go build -ldflags '-s -w -X main.Version=${VERSION}' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN}-darwin-${ARCH}-${VERSION}${BIN_MACOS} main.go
	@echo 'done'
