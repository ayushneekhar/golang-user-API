package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var client *mongo.Client

func CreateUserEndpoint(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type", "application/json")
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	hash, _ := HashPassword(user.Password)
	user.Password = hash
	collection := client.Database("instagram").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
}

func CreatePostEndpoint(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type", "application/json")
	var post Post
	json.NewDecoder(request.Body).Decode(&post)
	collection := client.Database("instagram").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, post)
	json.NewEncoder(response).Encode(result)
}

func GetUsersEndpoint(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type", "application/json")
	var users []User
	collection := client.Database("instagram").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx){
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(users)
}

func GetUserPostsEndpoint(response http.ResponseWriter, request *http.Request){
    response.Header().Add("content-type", "application/json")
    params := mux.Vars(request)
    id, _ := params["id"]
    var posts []Post
    collection := client.Database("instagram").Collection("posts")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, Post{UserID: id})
    if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{"message": "` + err.Error() + `"}`))
        return
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx){
		var post Post
        cursor.Decode(&post)
        posts = append(posts, post)
    }
	fmt.Println(posts)
    if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{"message": "` + err.Error() + `"}`))
        return
    }
    json.NewEncoder(response).Encode(posts)

}

func GetUserEndpoint(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	fmt.Println(id)
	var user User
	collection := client.Database("instagram").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, User{ID: id}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(user)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	router := mux.NewRouter()
	router.HandleFunc("/user", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/users", GetUsersEndpoint).Methods("GET")
	router.HandleFunc("/user/{id}", GetUserEndpoint).Methods("GET")

	router.HandleFunc("/post", CreatePostEndpoint).Methods("POST")
	router.HandleFunc("/posts/users/{id}", GetUserPostsEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)
}