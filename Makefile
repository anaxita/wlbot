win-build:
	GOOS=windows GOARCH=amd64 go build -o bin/wlbot-windows-amd64.exe cmd/main.go