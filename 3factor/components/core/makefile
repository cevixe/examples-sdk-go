.PHONY: build

build-CdcFn:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler cdc/cmd/main.go
build-StateFn:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler state/cmd/main.go
build-NotifyFn:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler notify/cmd/main.go
