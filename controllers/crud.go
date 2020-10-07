package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sanketgupta07/go-rest-example/util/header"

	"github.com/gorilla/mux"
	"github.com/sanketgupta07/go-rest-example/dbconfig"
	"github.com/sanketgupta07/go-rest-example/model"
	"go.mongodb.org/mongo-driver/bson"
)

type allEvents []model.User

var userCollection = dbconfig.GetDB().Database("goTest").Collection("users")

//HomeLink -- welcome page for browser
func HomeLink(w http.ResponseWriter, r *http.Request) {
	header.EnableCors(&w)
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "Golang + Angular Starter Kit",
	}

	json.NewEncoder(w).Encode(data)
}

//CreateUser a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	header.EnableCors(&w)
	w.Header().Set("Content-Type", "application/json") // for adding       //Content-type
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user) // storing in person   //variable of type user
	if err != nil {
		fmt.Print(err)
	}
	log.Println("User: ", user)
	insertResult, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Print(err)
	}
	log.Println("inserted User: ", insertResult)
	// json.NewEncoder(w).Encode(insertResult.InsertedID) // return the //mongodb ID of generated document
	var iUser model.User
	er := userCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: insertResult.InsertedID}}).Decode(&iUser)
	if er != nil {
		log.Print(er)
	}

	json.NewEncoder(w).Encode(iUser)
}

// GetOneUser from collection
func GetOneUser(w http.ResponseWriter, r *http.Request) {
	header.EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	// Get the ID from the url
	user := model.User{}
	name := mux.Vars(r)["name"] //in string
	log.Println("Name: ", name)
	err := userCollection.FindOne(context.TODO(), bson.D{{Key: "name", Value: name}}).Decode(&user)
	if err != nil {
		log.Print(err)
	}

	json.NewEncoder(w).Encode(user)
}

// GetAllUsers from collection
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	header.EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	var users []model.User

	cursor, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Print(err)
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var user model.User
		if err = cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

//UpdateUser  to update event
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	header.EnableCors(&w)
	w.Header().Set("Content-Type", "application/json") // for adding       //Content-type
	name := mux.Vars(r)["name"]                        //in string
	log.Println("Name: ", name)
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user) // storing in person   //variable of type user
	if err != nil {
		fmt.Print(err)
	}
	log.Println("User: ", user)
	filter := bson.D{{Key: "name", Value: name}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "address", Value: user.Address}}}}
	updatedResult, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Print(err)
	}
	log.Println("inserted User: ", updatedResult)
	json.NewEncoder(w).Encode(updatedResult) // return the //mongodb ID of generated document

}

//DeleteUser from collection
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	header.EnableCors(&w)
	w.Header().Set("Content-Type", "application/json") // for adding       //Content-type
	name := mux.Vars(r)["name"]                        //in string
	log.Println("Name: ", name)

	del, err := userCollection.DeleteOne(context.TODO(), bson.D{{Key: "name", Value: name}})
	if err != nil {
		log.Print(err)
	}
	log.Println("Deleted Result: ", del)
	json.NewEncoder(w).Encode(del) // return the //mongodb ID of generated document

}
