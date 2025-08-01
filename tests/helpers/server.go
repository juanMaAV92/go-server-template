package helpers

import (
	"github.com/juanMaAV92/go-server-template/cmd"
	"github.com/juanMaAV92/go-server-template/platform/config"
)

func NewTestServer() *cmd.Instance {
	cfg, _ := config.Load("local")
	testServer, _ := cmd.NewServer(cfg, nil)
	return testServer
}
