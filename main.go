package main

import (
	"context"
	"os"
	"os/signal"

	qmq "github.com/rqure/qmq/src"
)

func main() {
	ctx := context.Background()

	app := qmq.NewQMQApplication(ctx, "")
	app.Initialize(ctx)
	defer app.Deinitialize(ctx)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	for {
		select {
		case <-ctx.Done():
			return
		case <-sigint:
			return
		default:
			app.Logger().PrintNextEntry(ctx)
		}
	}
}
