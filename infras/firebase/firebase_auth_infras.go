package infras

import (
	"context"
	//"com.pegatech.faceswap/common/errors/error_app"
	"com.pegatech.faceswap/app_config"
	"com.pegatech.faceswap/infras/logger"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func FirebaseAuthClientProviderByConfig(ctx context.Context, FirebaseConfig app_config.FirebaseConfig) (*auth.Client, error) {
	opt := option.WithCredentialsFile(FirebaseConfig.ServiceKeyPath)
	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		logger.Fatal("Failed to create new Firebase App", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		logger.Fatal("Failed to connect to Firebase", err)
	}

	return client, nil
}