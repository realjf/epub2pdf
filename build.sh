#!/bin/bash
# ##############################################################################
# # File: build.sh                                                             #
# # Project: epub2pdf                                                          #
# # Created Date: 2023/09/10 20:08:17                                          #
# # Author: realjf                                                             #
# # -----                                                                      #
# # Last Modified: 2023/09/10 20:12:34                                         #
# # Modified By: realjf                                                        #
# # -----                                                                      #
# # Copyright (c) 2023 realjf                                                  #
# ##############################################################################

if test "$${GOOS}" = "linux"; then
    @env CGO_ENABLED=1 GOOS=linux GOARCH=$${ARCH} go build -ldflags '-s -w -X main.Version=$${VERSION}' -gcflags="all=-trimpath=$${PWD}" -asmflags="all=-trimpath=$${PWD}" -o $${BIN} main.go
elif test "$${GOOS}" = "windows"; then
    @env CGO_ENABLED=1 GOOS=windows GOARCH=$${ARCH} go build -ldflags '-s -w -X main.Version=$${VERSION}' -gcflags="all=-trimpath=$${PWD}" -asmflags="all=-trimpath=$${PWD}" -o $${BIN_WIN} main.go
elif test "$${GOOS}" = "darwin"; then
    @env CGO_ENABLED=1 GOOS=darwin GOARCH=$${ARCH} go build -ldflags '-s -w -X main.Version=$${VERSION}' -gcflags="all=-trimpath=$${PWD}" -asmflags="all=-trimpath=$${PWD}" -o $${BIN_MACOS} main.go
else
    echo "Target OS is not supported"
fi
echo 'done'
