.PHONY: build

build-CreateProductFn:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler product/cmd/create/main.go

build-UpdateProductFn:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(ARTIFACTS_DIR)/handler product/cmd/update/main.go
