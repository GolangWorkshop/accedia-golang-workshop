package util

import (
	"context"
	"os"
	"os/signal"

	baseLog "github.com/rs/zerolog/log"
)

var log = baseLog.With().Str("package", "util").Logger()

func GetContext() context.Context {

	ctx := context.Background()

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		select {
		case <-c:
			log.Info().Msg("received interrupt signal, terminating context")
			cancel()
		case <-ctx.Done():
		}
	}()

	return ctx
}
