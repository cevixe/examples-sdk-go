package main

import (
	"github.com/cevixe/aws-sdk-go/runtime"
	"github.com/cevixe/examples-sdk-go/3factor/services/product/pkg/application"
)

func main() {
	app := application.New()
	runtime.Start(app.UpdateProduct)
}
