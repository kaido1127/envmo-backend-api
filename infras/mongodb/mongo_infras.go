package infras

import (
	"context"

	"com.pegatech.faceswap/app_config"
	"com.pegatech.faceswap/infras/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoDBProviderByConfig(mongodbConfig app_config.MongoDbConfig) *mongo.Database {
	clientOptions := options.Client().ApplyURI(mongodbConfig.ConnectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB", err)
	}
	return client.Database(mongodbConfig.DbName)
}

func PaginationOptionByPageAndLimit(page int, limit int) *options.FindOptions {
	pageInt64 := int64(page)
	limitInt64 := int64(limit)
	skip := (pageInt64 - 1) * limitInt64

	options := options.Find()
	options.SetLimit(limitInt64)
	options.SetSkip(skip)
	return options
}

// func GetOrCreateCollection(db *mongo.Database, collectionName string) *mongo.Collection {
// 	ctx := context.Background()
// 	fmt.Println("[MongoDB_Infras] Start getlist collection names of", db.Name())
// 	collectionNames, err := db.ListCollectionNames(ctx, nil)
// 	if err != nil {
// 		fmt.Println("[MongoDB_Infras] Failed to get list collection names", err)

// 		return nil
// 	}
// 	fmt.Println("[MongoDB_Infras] Finish get list collection names")

// 	collectionExists := false
// 	for _, name := range collectionNames {
// 		if name == collectionName {
// 			collectionExists = true
// 			break
// 		}
// 	}
// 	if !collectionExists {
// 		fmt.Println("Create collection", collectionName)
// 		err := db.CreateCollection(ctx, collectionName)
// 		if err != nil {
// 			fmt.Println("Create collection error")
// 			return nil
// 		}
// 	}
// 	return db.Collection(collectionName)
// }
