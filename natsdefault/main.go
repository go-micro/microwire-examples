package main

import (
	"github.com/go-micro/microwire"
	mWire "github.com/go-micro/microwire/wire"
	"go-micro.dev/v4/logger"

	_ "github.com/go-micro/microwire/plugins/cli/urfave"
	_ "github.com/go-micro/plugins/v4/broker/nats"
	_ "github.com/go-micro/plugins/v4/registry/nats"
	_ "github.com/go-micro/plugins/v4/transport/nats"
)

func main() {
	service, err := microwire.DefaultApp(
		mWire.Name("natsdefault"),
		mWire.Usage("A POC for go-micro.dev/v5"),
		mWire.Version("v0.0.1"),
		mWire.ArgPrefix(""),

		mWire.Component(mWire.ComponentBroker, "nats"),
		mWire.Component(mWire.ComponentRegistry, "nats"),
		mWire.Component(mWire.ComponentTransport, "nats"),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
