package main

import (
	"functions/product/pkg/application"
	"github.com/cevixe/aws-sdk-go/runtime"
)

func main() {
	app := application.New()
	runtime.Start(app.UpdateProduct)
}
