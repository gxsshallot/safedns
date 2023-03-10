.PHONY: build
build:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags "-w -s" -o dist/safedns-windows.exe cmd/main.go
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags "-w -s" -o dist/safedns-mac cmd/main.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags "-w -s" -o dist/safedns-linux cmd/main.go
