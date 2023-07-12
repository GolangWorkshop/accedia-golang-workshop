package webserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/GolangWorkshop/library/store"
	baseLog "github.com/rs/zerolog/log"
)

var log = baseLog.With().Str("package", "webserver").Logger()

func Start(ctx context.Context, host string, port string, store store.Store) error {

	sm := http.NewServeMux()
	addRoutes(sm, store)

	smWithMiddlewares := UserContextMiddleware(HTTPLogMiddleware(sm))

	addr := fmt.Sprintf("%s:%s", host, port)
	log.Info().Msgf("server listening on %s", addr)

	s := &http.Server{
		Addr:    addr,
		Handler: smWithMiddlewares,
	}

	go func() {
		// wait for context to finish
		// shut server down gracefully
		<-ctx.Done()
		ctxBg := context.Background()
		ctxWithDeadline, cancel := context.WithDeadline(ctxBg, time.Now().Add(time.Second*3))

		log.Info().Msg("server shutting down gracefully, will force in 3s")
		cancel()
		s.Shutdown(ctxWithDeadline)
	}()

	err := s.ListenAndServe()

	log.Error().Err(err).Msg("web server failed")
	return err
}
