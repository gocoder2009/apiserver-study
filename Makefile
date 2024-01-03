SHELL := /bin/bash
BASEDIR = $(shell pwd)

# build with version infos
# 注意其中[-ldflags -X importpath.name=value](https://golang.org/cmd/link)
versionDir = "apiserver/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

# -w 为去掉调试信息（无法使用 gdb 调试），这样可以使编译后的二进制文件更小。
ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

all: go.tool build
	@echo "make all finished."
build:
	@go build -v -ldflags ${ldflags} .
clean:
	rm -f apiserver
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
go.tool:
	gofmt -w .
	go vet .
ca:
	openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"

help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make go.tool - run go tool 'fmt' and 'vet'"
	@echo "make ca - generate ca files"

.PHONY: build clean go.tool ca help


