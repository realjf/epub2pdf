BIN="bin/epub2pdf"
BIN_WIN="bin/epub2pdf.exe"
BIN_MACOS="bin/epub2pdf.app"

build:
	@env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN} main.go
	@echo 'done'

build_win:
	@env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags '-s -w' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN_WIN} main.go
	@echo 'done'

build_darwin:
	@env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN_MACOS} main.go
	@echo 'done'
