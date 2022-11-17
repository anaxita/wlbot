win-build:
	VERSION=$(shell git describe --tags --always) GOOS=windows GOARCH=amd64 go build -o bin/wlbot-$(shell git describe --tags --always)-windows-amd64.exe cmd/main.go