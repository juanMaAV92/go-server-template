// Package main provides the entry point for the Go Server Template application
//
// @title Go Server Template API
// @version 1.0.0
// @description API documentation for Go Server Template microservice
//
// @contact.name juanMaAV92
// @contact.url https://github.com/juanMaAV92/go-server-template

// @host localhost:8080
// @BasePath /go-server-template
// @schemes http https
package main

import "github.com/juanMaAV92/go-server-template/cmd"

func main() {
	cmd.Start()
}
