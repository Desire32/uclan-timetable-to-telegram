package timetable

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
	"timetable/internal/services/timetable/data"
)

type MongoService struct {
}

func (m *MongoService) MongoConnect() (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGODB_URI")).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (m *MongoService) MongoSend(jsonData string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := m.MongoConnect()
	if err != nil {
		return fmt.Errorf("Connection error:  %w", err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, context.TODO())
	if err := json.Unmarshal([]byte(jsonData), &data.Schedule); err != nil {
		return fmt.Errorf("parsing error : %w", err)
	}
	collection := client.Database("Timetable").Collection("schedule")

	for _, entry := range data.Schedule {
		_, err := collection.InsertOne(ctx, entry)
		if err != nil {
			log.Printf("Error day loading %s: %v", entry.Day, err)
		} else {
			log.Printf("Success day loading %s", entry.Day)
		}
	}
	return nil
}
