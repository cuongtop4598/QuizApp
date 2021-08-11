package db

import (
	"context"
	"time"

	"github.com/cuongtop4598/QuizWithGo/QuizApp/pkg/setting"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Resource struct {
	DB *mongo.Database
}

// Close use this method to close database connection
func (r *Resource) Close() {
	logrus.Warning("closing all db connections")
}
func InitResource() (*Resource, error) {

	host := setting.DatabaseSetting.Host
	dbName := setting.DatabaseSetting.DBName
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(host))

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = mongoClient.Connect(ctx)
	if err != nil {
		return nil, err
	}
	mongoClient.Ping(ctx, nil)
	return &Resource{DB: mongoClient.Database(dbName)}, nil
}
