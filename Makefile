VERSION = 0.0.1
PACKAGE = github.com/nuxy/go-webview-src-builder
DEBUG   = $(shell echo $(MAKEFLAGS) | grep -q -- "--debug" && echo true || echo false)
LDFLAGS = "-X main.Version=$(VERSION) -X main.DevTools=$(DEBUG)"

ifeq ($(OS), Windows_NT)
    LDFLAGS := "-H windowsgui"
endif

run:
	go run $(GOFLAGS) -ldflags $(LDFLAGS) ./src.go

build:
	go build -x $(GOFLAGS) -ldflags $(LDFLAGS) -o ./bin/webview-src $(PACKAGE)

build-darwin:
	GOOS=darwin GOARCH=amd64 go build $(GOFLAGS) -ldflags $(LDFLAGS) -o ./bin/webview-src-$(VERSION)-osx-64 $(PACKAGE)

build-linux:
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -ldflags $(LDFLAGS) -o ./bin/webview-src-$(VERSION)-linux-64 $(PACKAGE)

build-windows:
	GOOS=windows GOARCH=amd64 go build $(GOFLAGS) -ldflags $(LDFLAGS) -o ./bin/webview-src-$(VERSION)-windows-64.exe $(PACKAGE)

install:
	go install -x $(GOFLAGS) -ldflags $(LDFLAGS) $(PACKAGE)
