package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("failed to connect mongo: %s\n", err.Error())
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("failed to ping mongo: %s\n", err.Error())
	}
	data := Data{X: "0", Z: "9"}
	// data := Data{Y: "555"}
	internalData := InternalData{X: data.X, Y: data.Y, Z: data.Z}

	id, _ := primitive.ObjectIDFromHex("6628b7797075760792f31156")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": internalData}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := client.Database("test").Collection("test").FindOneAndUpdate(ctx, filter, update, opts)
	fmt.Println(result.Raw())
}

// request protobuf
type Data struct {
	X string `bson:"xx"`
	Y string `bson:"yy"`
	Z string `bson:"zz"`
}

// domains
type InternalData struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	X  string             `bson:"xx,omitempty"`
	Y  string             `bson:"yy,omitempty"`
	Z  string             `bson:"zz,omitempty"`
}
