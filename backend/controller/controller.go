package controller

//This package is meant to be a controller to our database, and initializes our connection.
import (
	"context"
	"fmt"
	"log"

	"github.com/Tarunss/food-blog/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// how we connect to our server
const connectionString = "mongodb+srv://tarunsohal:superstrongpassword6969@cluster0.muyxchk.mongodb.net/"
const dbName = "BlogPosts"
const colName = "postList"

// creating a connection
var collection *mongo.Collection

//connect with mongoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success!")
	collection = client.Database(dbName).Collection(colName)

	//if collection instance is ready
	fmt.Println("Colletion reference is ready")
}

// MongoDB helper method
func insertOnePost(post model.BlogPost) {
	inserted, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in DB with ID:", inserted.InsertedID)
}
