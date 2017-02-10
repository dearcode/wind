package api

import (
	"github.com/dearcode/petrel/handler"
)

func init() {
	handler.Server.AddInterface(&sites{}, "")
}
