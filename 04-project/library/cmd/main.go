package main

import (
	"github.com/GolangWorkshop/library/clients"
	"github.com/GolangWorkshop/library/config"
	"github.com/GolangWorkshop/library/store"
	"github.com/GolangWorkshop/library/util"
	"github.com/GolangWorkshop/library/webserver"
	baseLog "github.com/rs/zerolog/log"
)

func main() {
	log := baseLog.With().Str("package", "main").Logger()

	ctx := util.GetContext()

	log.Info().Msg("loading env vars from .env file")
	util.LoadEnvFromFile(".env")

	log.Info().Msg("gotting config based on env vars")
	cfg := config.GetConfig()

	log.Info().Msgf("connecting to db %s", cfg.DatabaseConfig.ConnectionString)
	db, err := clients.NewPgClient(*cfg.DatabaseConfig)
	if err != nil {
		log.Error().Err(err).Msg("could not connect to db, shutting down")
		return
	}
	defer func() {
		db.Close()
		log.Info().Msg("disconnected from db")
	}()

	log.Info().Msg("connected to db")

	store := store.NewStore(db)

	log.Info().Msg("starting server")
	err = webserver.Start(ctx, cfg.WebServerConfig.Host, cfg.WebServerConfig.Port, store)
	log.Error().Err(err).Msg("shutting down")
}
