package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDBClient struct {
	instrumentsCollection *mongo.Collection
}

type instrument struct {
	Id             int32  `json:"instrumentId"`
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	InstrumentType string `json:"instrumentType"`
}

func newMongoDBClient() *mongoDBClient {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	instrumentsCollection := client.Database("market").Collection("instruments")
	fmt.Println("Connected to MongoDB!")
	return &mongoDBClient{instrumentsCollection}
}

func (m *mongoDBClient) getAllInstruments() ([]bson.M, error) {
	var results []bson.M
	cur, err := m.instrumentsCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cur.All(context.Background(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (m *mongoDBClient) addInstrument(data instrument) error {
	obj := bson.M{
		"instrumentId":   data.Id,
		"name":           data.Name,
		"symbol":         data.Symbol,
		"instrumentType": data.InstrumentType,
	}
	if _, err := m.instrumentsCollection.InsertOne(context.TODO(), obj); err != nil {
		return err
	}

	return nil
}

func (m *mongoDBClient) deleteInstrument(instrumentID int) error {
	_, err := m.instrumentsCollection.DeleteOne(context.TODO(), bson.M{"instrumentId": instrumentID})
	if err != nil {
		return err
	}
	return nil
}
