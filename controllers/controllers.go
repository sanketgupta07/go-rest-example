package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sanketgupta07/go-rest-example/model"

	"github.com/gorilla/mux"
)

type allEvents []model.Event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

//HomeLink -- welcome page for browser
func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

//CreateEvent a new event
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var newEvent []model.Event
	err = json.Unmarshal(body, &newEvent)
	if err != nil {
		panic(err)
	}

	// Add the newly created event to the array of events
	events = append(events, newEvent...)

	// Return the 201 created status code
	w.WriteHeader(http.StatusCreated)
	// Return the newly created event
	json.NewEncoder(w).Encode(newEvent)
}

func GetOneEvent(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	eventID := mux.Vars(r)["id"]

	// Get the details from an existing event
	// Use the blank identifier to avoid creating a value that will not be used
	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	eventID := mux.Vars(r)["id"]
	var updatedEvent []model.Event
	// Convert r.Body into a readable formart
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	log.Println(string(reqBody))
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent[0].Title
			singleEvent.Description = updatedEvent[0].Description
			events[i] = singleEvent
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	eventID := mux.Vars(r)["id"]

	// Get the details from an existing event
	// Use the blank identifier to avoid creating a value that will not be used
	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}
