.PHONY:build
build:
	GOARCH=arm64 GOOS=android go build
