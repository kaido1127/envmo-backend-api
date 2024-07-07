package datasource

import (
	"context"
	"envmo/module/env_info/data/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type envInfoMongoDataSourceImpl struct {
	collection mongo.Collection
}

// GetHistoryByDeviceID implements EnvInfoMongoDataSource.
func (e envInfoMongoDataSourceImpl) GetHistoryByDeviceID(ctx context.Context, deviceID string, startTime int64, endTime int64) ([]model.HistoryEnvInfoModel, error) {

	filter := bson.M{
		"device_id": deviceID,
		"created_at": bson.M{
			"$gte": startTime,
			"$lte": endTime,
		},
	}

	opts := options.Find().SetSort(bson.M{"created_at": -1})
	cur, err := e.collection.Find(ctx, filter, opts)

	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var records []model.HistoryEnvInfoModel
	for cur.Next(ctx) {
		var record model.HistoryEnvInfoModel

		if cur.Decode(&record) == nil {
			records = append(records, record)
		}
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

// Insert implements EnvInfoMongoDataSource.
func (e envInfoMongoDataSourceImpl) InsertMany(ctx context.Context, historyEnvInfoModel []model.HistoryEnvInfoModel) (int, error) {
	var interfaceSlice []interface{}
	for _, info := range historyEnvInfoModel {
		interfaceSlice = append(interfaceSlice, info)
	}

	result, err := e.collection.InsertMany(ctx, interfaceSlice)
	if err != nil {
		return 0, err
	}

	return len(result.InsertedIDs), nil
}

func NewEnvInfoMongoDataSource(mongoDb mongo.Database) EnvInfoMongoDataSource {
	return envInfoMongoDataSourceImpl{
		collection: *mongoDb.Collection(model.EnvInfoMongoCollectionName),
	}
}
