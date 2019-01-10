package main

import (
	"github.com/travelgateX/go-io/log"
	logconf "hub-aggregator/common/config/log"
	"hub-aggregator/common/kit/routing"
	"rfc/presenters/pkg/presenter/config"
)

func main() {
	logger := log.NewStdLogger()
	logb := logconf.NewBuilder(logconf.SetLogger(logger))
	server := routing.NewServer(":8080", config.NewRoutes(*logb))
	logger.Info("running")
	server.ListenAndServe()
}
