.PHONY: all build run gotool clean help

all:
	gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./main.go ./config.json

gotool:
	go fmt ./
	go vet ./

clean:
	@if	[ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 GO 代码，并编译生成二进制文件"
