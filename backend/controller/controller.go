package controller

//This package is meant to be a controller to our database, and initializes our connection.
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Tarunss/food-blog/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// how we connect to our server
const dbName = "BlogPosts"
const colName = "postList"
const authName = "auth"

var secret = os.Getenv("JWT_SECRET")

// creating a collection
var collection *mongo.Collection
var auth *mongo.Collection

//connect with mongoDB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := os.Getenv("MONGODB_SECRET")
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success!")
	collection = client.Database(dbName).Collection(colName)
	auth = client.Database(dbName).Collection(authName)

	//if collection instance is ready
	fmt.Println("Collection reference is ready")
}

// MongoDB helper methods

func registerUser(cred model.UsePass) {
	// First, we hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cred.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error hashing password:", err)
	}

	// Replace plaintext password with hashed password
	cred.Password = string(hashedPassword)
	inserted, err := auth.InsertOne(context.Background(), cred)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Registered 1 user within the DB", inserted.InsertedID)
}

// getUser fetches a single user document by _id
func getUser(username string) model.UsePass {
	// MongoDB filter
	filter := bson.M{"username": username}

	// Perform the query and decode into model.UsePass
	var user model.UsePass
	err := auth.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Fatal("Error finding user:", err)
	}

	// For debugging
	fmt.Println("Fetched user:", user)

	return user
}

// Insert One Post
func insertOnePost(post model.BlogPost) {
	inserted, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 post in DB with ID:", inserted.InsertedID)
}

// Update One Post
func updateOnePost(postID string, body string, title string, summary string) {
	id, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"body": body, "title": title, "summary": summary}}

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	//create our post model
	var post model.BlogPost
	_ = json.NewDecoder(r.Body).Decode(&post)
	insertOnePost(post)

	json.NewEncoder(w).Encode(post)
}

// Update Post
func UpdateOnePost(w http.ResponseWriter, r *http.Request) {
	//enableCors(&w)
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	//Decode "body" from updated post
	var post model.BlogPost
	_ = json.NewDecoder(r.Body).Decode(&post)
	//get our pararmeters from mux router (this is to get our ID)
	params := mux.Vars(r)
	//fmt.Println(post.Body)
	//fmt.Println(params["id"])
	updateOnePost(params["id"], post.Body, post.Title, post.Summary)
	// send back a json response of id updated, and update
	encoder := json.NewEncoder(w)
	encoder.Encode(params["id"])
	encoder.Encode(post.Summary)
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

// Delete all posts
func DeleteAllPosts(w http.ResponseWriter, r *http.Request) {
	//enableCors(&w)
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllPosts()
	//send back a json representation of posts deleted
	json.NewEncoder(w).Encode(count)
}

/**
------------------------------------------------AUTH FUNCTIONS -------------------------------------------------------------------
**/
// Replace this with a secure secret stored in .env!
var jwtSecret = []byte(secret)

// GenerateJWT creates a JWT for a given username
func GenerateJWT(username string) (string, error) {
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // token expires in 24 hours
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// POST request when user submits password
func PostToken(w http.ResponseWriter, r *http.Request) {
	// set our headers
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	// Decode the post request
	var usepass model.UsePass
	_ = json.NewDecoder(r.Body).Decode(&usepass)
	// fetch from the database with user, check encrypted password (using bcrypt), then return a token or a failure

	userFromDB := getUser(usepass.Username)
	err := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(usepass.Password))
	if err != nil {
		// Password doesn’t match
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	token, err := GenerateJWT(userFromDB.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour), // expires in 1 day
		HttpOnly: true,                           // prevent JS access
		Secure:   true,                           // send only over HTTPS
		SameSite: http.SameSiteLaxMode,           // adjust based on your CSRF policy
		Path:     "/",                            // cookie path
	})

	// Send success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
	})
}

// POST request when user registers
func Register(w http.ResponseWriter, r *http.Request) {
	//enableCors(&w)
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	//create our post model
	var cred model.UsePass
	_ = json.NewDecoder(r.Body).Decode(&cred)
	registerUser(cred)

	json.NewEncoder(w).Encode(cred)
}
