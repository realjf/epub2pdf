VERSION="0.0.3"
BIN="bin/epub2pdf"
BIN_WIN=".exe"
BIN_MACOS=""
ARCH="amd64"

.PHONY: build build_win build_darwin tag dtag


build:
	@env CGO_ENABLED=1 GOOS=linux GOARCH=${ARCH} go build -ldflags '-s -w -X main.Version=${VERSION}' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN}-linux-${ARCH}-${VERSION} main.go
	@echo 'done'

build_win:
	@env CGO_ENABLED=1 GOOS=windows GOARCH=${ARCH} go build -ldflags '-s -w -X main.Version=${VERSION}' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN}-windows-${ARCH}-${VERSION}${BIN_WIN} main.go
	@echo 'done'

build_darwin:
	@env CGO_ENABLED=1 GOOS=darwin GOARCH=${ARCH} go build -ldflags '-s -w -X main.Version=${VERSION}' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN}-darwin-${ARCH}-${VERSION}${BIN_MACOS} main.go
	@echo 'done'

push:
	@git add -A && git commit -m "update" && git push origin master


# make tag t=<your_version>
tag:
	@echo '${t}'
	@git tag -a ${t} -m "${t}" && git push origin ${t}

dtag:
	@echo 'delete ${t}'
	@git push --delete origin ${t} && git tag -d ${t}
