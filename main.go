package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

const itchURL = "https://itch.io/api/1/jwt/me"

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	if err := initializer.RegisterBeforeAuthenticateCustom(BeforeAuthItch); err != nil {
		logger.Error("Unable to register itch auth module: %v", err)
		return err
	}
	logger.Info("Itch auth module loaded")
	return nil
}

func BeforeAuthItch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateCustomRequest) (*api.AuthenticateCustomRequest, error) {
	token := in.GetAccount().GetId()

	if token == "" {
		return in, errors.New("Missing itch.io token")
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", itchURL, nil)
	if err != nil {
		return in, fmt.Errorf("http.NewRequest failed: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return in, fmt.Errorf("client.Do failed: %v", err)
	}
	if resp.StatusCode >= 300 {
		return in, fmt.Errorf("itch request error: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	var response struct {
		User *struct {
			Username    string
			DisplayName string `json:"display_name"`
			ID          int
		}
		Errors []string
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return in, fmt.Errorf("Failed to unmarshal itch response: %v", err)
	}
	if len(response.Errors) > 0 {
		return in, fmt.Errorf("itch response has errors: %v", response.Errors)
	}

	// Add a prefix to ensure our ID is >= 6 characters (itch IDs may be shorter)
	in.Account.Id = fmt.Sprintf("itch-%d", response.User.ID)

	if in.Username == "" {
		in.Username = response.User.Username
	}

	return in, nil
}
