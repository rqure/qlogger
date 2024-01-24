package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"time"

	qmq "github.com/rqure/qmq/src"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	ctx := context.Background()

	app := qmq.NewQMQApplication(ctx, "")
	app.Initialize(ctx)
	defer app.Deinitialize(ctx)
}
