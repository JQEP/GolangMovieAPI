package database

import (
	"context"
	"fmt"
	"go-graphql-mongodb-api/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

//Connect initializes the mongo.Client by starting background monitoring goroutines to monitor the deployment state
func Connect(dbUrl string) *DB {
	//NewClient generates a new client to connect to deployment indicated by the URI which is the mongo.Client
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//While monitoring we use client.Ping to verify the connection to the mongo.Client was created successfully
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	//If successful connection is initiated, we return the DB struct with the mongodb connection
	return &DB{
		client: client,
	}
}

func (db *DB) InsertMovieByID(movie model.NewMovie) *model.Movie {
	movieColl := db.client.Database("graphql-mongodb-api-db").Collection("movie")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	inserg, err := movieColl.InsertOne(ctx, bson.M{"name": movie.Name, "description": movie.Description})
	if err != nil {
		log.Fatal(err)
	}

	insertedId := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnMovie := model.Movie{ID: insertedId, Name: movie.Name, Description: movie.Description}
	return &returnMovie
}

//FindMovieById is based on getting a single result from the database.
func (db *DB) FindMovieById(id string) *model.Movie {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	movieColl := db.client.Database("graphql-mongodb-api-db").Collection("movie")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := movieColl.FindOne(ctx, bson.M{"_id": ObjectID})
	movie := model.Movie{ID: id}

	res.Decode(&movie)

	return &movie
}

//All will get all the movie lists saved to the movie collection.
func (db *DB) All() []*model.Movie {
	movieColl := db.client.Database("graphql-mongodb-api-db").Collection("movie")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := movieColl.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var movies []*model.Movie
	for cur.Next(ctx) {
		sus, err := cur.Current.Elements()
		fmt.Println(sus)
		if err != nil {
			log.Fatal(err)
		}
		movie := model.Movie{ID: (sus[0].String()), Name: (sus[1].String()), Description: (sus[2].String())}

		movies = append(movies, &movie)

	}
	return movies
}
