package controller

//This package is meant to be a controller to our database, and initializes our connection.
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Tarunss/food-blog/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	fmt.Println("Collection reference is ready")
}

// MongoDB helper methods
// Insert One Post
func insertOnePost(post model.BlogPost) {
	inserted, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in DB with ID:", inserted.InsertedID)
}

// Update One Post
func updateOnePost(postID string, body string, title string) {
	id, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"body": body, "title": title}}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", res.ModifiedCount)
}

// Delete 1 record

func deleteOnePost(postID string) {
	id, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted count:", res.DeletedCount)
}

// Delete all posts
func deleteAllPosts() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Amount of posts deleted:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get All posts
func getAllPosts() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var posts []primitive.M
	for cur.Next(context.Background()) {
		var post bson.M
		err := cur.Decode(&post)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}
	defer cur.Close(context.Background())
	return posts
}

// Actual controllers - file
// enabling cors for this origin
// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }

// get all posts
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	//enableCors(&w)
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allPosts := getAllPosts()
	json.NewEncoder(w).Encode(allPosts)
}

// Create Post

func InsertOnePost(w http.ResponseWriter, r *http.Request) {
	//enableCors(&w)
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	//create our post model
	var post model.BlogPost
	//decode our recieved post from json
	_ = json.NewDecoder(r.Body).Decode(&post)
	insertOnePost(post)
	//send back json response of post inserted
	json.NewEncoder(w).Encode(post)
}

// Update Post
func UpdateOnePost(w http.ResponseWriter, r *http.Request) {
	//enableCors(&w)
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	//Decode "body" from updated post
	var post model.BlogPost
	_ = json.NewDecoder(r.Body).Decode(&post)
	//get our pararmeters from mux router (this is to get our ID)
	params := mux.Vars(r)
	//fmt.Println(post.Body)
	//fmt.Println(params["id"])
	updateOnePost(params["id"], post.Body, post.Title)
	// send back a json response of id updated, and update
	encoder := json.NewEncoder(w)
	encoder.Encode(params["id"])
	encoder.Encode(post.Body)
}

// Delete a post
func DeleteOnePost(w http.ResponseWriter, r *http.Request) {
	//enableCors(&w)
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOnePost(params["id"])
	//send back a json response of id
	json.NewEncoder(w).Encode(params["id"])
}

// delete All posts
func DeleteAllPosts(w http.ResponseWriter, r *http.Request) {
	//enableCors(&w)
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllPosts()
	//send back a json representation of posts deleted
	json.NewEncoder(w).Encode(count)
}
