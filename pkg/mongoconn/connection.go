package mongoconn

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DatabaseService struct {
	client       *mongo.Client
	DatabaseName string
	Database     *mongo.Database
}

var (
	databaseService *DatabaseService
)

func (service *DatabaseService) Client() *mongo.Client {
	return service.client
}

func (service *DatabaseService) newConnection(databaseName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// TODO: Insert URI MONGODB
	uri := ""
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal()
	}

	service.client = client

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalln("Ping mongoDB server error : ", err)
		return err
	}

	service.Database = service.client.Database(databaseName)

	log.Println("Mongo connected.")

	return nil
}

func (service *DatabaseService) NewService() error {
	if service.DatabaseName == "" {
		return errors.New("database name is required")
	}

	// Already connected
	if service.client != nil && service.Database != nil {
		return nil
	}

	if err := service.newConnection(service.DatabaseName); err != nil {
		return err
	}

	// set cache
	service.setCache()

	return nil
}

func (service *DatabaseService) setCache() {
	databaseService = service
}

func (service *DatabaseService) GetCache() *DatabaseService {
	return databaseService
}
