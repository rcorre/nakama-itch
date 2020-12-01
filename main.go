package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	if err := initializer.RegisterBeforeAuthenticateCustom(BeforeAuthItch); err != nil {
		logger.Error("Unable to register itch auth module: %v", err)
		return err
	}
	logger.Info("Itch auth module loaded")
	return nil
}

func BeforeAuthItch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateCustomRequest) (*api.AuthenticateCustomRequest, error) {
	if in.GetAccount().GetId() == "blahblah" {
		return in, nil
	}
	return in, errors.New("Bad name")
}
