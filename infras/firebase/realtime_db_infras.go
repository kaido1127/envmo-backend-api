package infras

import (
	"context"
	"os"

	"envmo/app_config"
	"envmo/infras/logger"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

func CreateRealtimeDbClientByConfig(realtimeDbCfg app_config.FirebaseConfig) *db.Client {
	home, err := os.Getwd()
	if err != nil {
		logger.Fatal("Failed to connect to RealtimeDB", err)
	}
	ctx := context.Background()
	opt := option.WithCredentialsFile(home + realtimeDbCfg.ServiceKeyPath)
	config := &firebase.Config{DatabaseURL: realtimeDbCfg.DatabaseUrl}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		logger.Fatal("Failed to connect to RealtimeDB", err)
	}
	client, err := app.Database(ctx)
	if err != nil {
		logger.Fatal("Failed to connect to RealtimeDB", err)
	}

	return client
}
