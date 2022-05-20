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

func (db *DB) InsertCourseByID(input model.NewCourse) *model.Course {
	courseColl := db.client.Database("graphql-mongodb-api-db").Collection("course")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	inserg, err := courseColl.InsertOne(ctx, bson.M{"name": input.Name, "subject": input.Subject, "hasinstructor": input.InstructorID})
	if err != nil {
		log.Fatal(err)
	}

	insertedId := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnMovie := model.Course{ID: insertedId, Name: input.Name, Subject: input.Subject, Instructor: &model.Instructor{ID: input.InstructorID}}
	return &returnMovie
}

//FindMovieById is based on getting a single result from the database.
func (db *DB) FindCourseById(id string) *model.Course {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	courseColl := db.client.Database("graphql-mongodb-api-db").Collection("course")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := courseColl.FindOne(ctx, bson.M{"_id": ObjectID})
	course := model.Course{ID: id}

	res.Decode(&course)

	return &course
}

//All will get all the movie lists saved to the movie collection.
func (db *DB) All() []*model.Course {
	movieColl := db.client.Database("graphql-mongodb-api-db").Collection("course")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := movieColl.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var courses []*model.Course
	for cur.Next(ctx) {
		sus, err := cur.Current.Elements()
		fmt.Println(sus)
		if err != nil {
			log.Fatal(err)
		}
		course := model.Course{ID: (sus[0].String()), Name: (sus[1].String()), Subject: (sus[2].String())}

		courses = append(courses, &course)

	}
	return courses
}

func (db *DB) FindInstructorById(id string) *model.Instructor {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	courseColl := db.client.Database("graphql-mongodb-api-db").Collection("course")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := courseColl.FindOne(ctx, bson.M{"_id": ObjectID})
	instructor := model.Instructor{ID: id}

	res.Decode(&instructor)

	return &instructor
}

func (db *DB) InsertInstructor(input *model.AddInstructor) *model.Instructor {
	courseColl := db.client.Database("graphql-mongodb-api-db").Collection("course")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	inserg, err := courseColl.InsertOne(ctx, bson.M{"firstname": input.Firstname, "lastname": input.Lastname, "salary": input.Salary})
	if err != nil {
		log.Fatal(err)
	}
	if len(input.ID) != 0 {
	}
	insertedId := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnMovie := model.Instructor{ID: insertedId, Firstname: input.Firstname, Lastname: input.Lastname, Salary: input.Salary}
	return &returnMovie
}
