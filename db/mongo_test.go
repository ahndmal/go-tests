package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"testing"
	"time"
)

type Workout struct {
	Id           string    `json:"_id" bson:"_id,omitempty"`
	Record       int64     `json:"record"`
	Sets         int       `json:"sets"`
	Comments     string    `json:"comments"`
	CreationDate time.Time `json:"creation_date" bson:"creation_date"`
	WorkoutDate  string    `json:"workout_date" bson:"workout_date"`
	Day          string    `json:"day"`
	Week         int       `json:"week"`
	WorkoutType  string    `json:"workout_type" bson:"workout_type"`
	Month        string    `json:"month"`
	Year         int       `json:"year"`
}

type WorkoutDto struct {
	Id           string `json:"_id" bson:"_id,omitempty"`
	Record       int64  `json:"record"`
	Sets         int    `json:"sets"`
	Comments     string `json:"comments"`
	CreationDate string `json:"creation_date" bson:"creation_date"`
	WorkoutDate  string `json:"workout_date" bson:"workout_date"`
	Day          string `json:"day"`
	Week         int    `json:"week"`
	WorkoutType  string `json:"workout_type" bson:"workout_type"`
	Month        string `json:"month"`
	Year         int    `json:"year"`
}

func TestGetCats(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://172.17.0.2:27017/cats"))
	if err != nil {
		log.Printf("Error connecting to MongoDB: %v", err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("cats").Collection("cats")

	var result bson.M

	coll.FindOne(context.TODO(), bson.D{{"age", 2}}).Decode(&result)

	println(result)
}

func TestGetWorkouts(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoUri := os.Getenv("MONGO_URL")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Printf("> error connecting to MongoDB: %v", err)
	}

	collection := client.Database("workouts").Collection("workouts")

	filter := bson.D{{"workout_date", "2023-11-05"}}

	curr, err := collection.Find(context.Background(), filter)
	defer curr.Close(ctx)

	if err != nil {
		log.Printf("> error dinding data in MongoDB: %v", err)
	}

	workouts := make([]Workout, 0)

	curr.All(context.Background(), &workouts)

	for _, w := range workouts {
		log.Println(w)
	}

}
