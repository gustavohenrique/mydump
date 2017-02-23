win:
	GOOS=windows GOARCH=amd64 go build -o mydump-windows-amd64.exe main.go

linux:
	GOOS=linux GOARCH=amd64 go build -o mydump-linux-amd64 main.go

mac:
	go build -o mydump-darwin-amd64
