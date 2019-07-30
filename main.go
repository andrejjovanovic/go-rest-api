package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Event struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Place     string `json:"place"`
	Speaker   string `json:"speaker"`
	EventType string `json:"eventType"`
	DateTime  string `json:"dateTime"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var EventList []Event

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// creates a new instance of a mux router
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/", homePage)
	// add our articles route and map it to our
	// returnAllArticles function like so
	muxRouter.HandleFunc("/getAllEvents", returnAllEvents)
	muxRouter.HandleFunc("/getEvent/{Id}", returnSingleEvent)
	muxRouter.HandleFunc("/createEvent", createNewEvent).Methods("POST")
	muxRouter.HandleFunc("/deleteEvent/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", muxRouter))
}


func returnAllEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllEvents")
	json.NewEncoder(w).Encode(EventList)
}

func returnSingleEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Id"]

	//fmt.Fprintf(w, "Key: " + key)
	for _, events := range EventList {
		if events.Id == key {
			json.NewEncoder(w).Encode(events)
		}
	}
}

func createNewEvent(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, "%+v", string(reqBody))
	var event Event
	json.Unmarshal(reqBody, &event)
	fmt.Fprintf(w, "%+v", string(reqBody))
	EventList = append(EventList, event)

	json.NewEncoder(w).Encode(event)
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]
	fmt.Println("Hello")

	// we then need to loop through all our articles
	for index, Event := range EventList {
		// if our id path parameter matches one of our
		// articles
		if Event.Id == id {
			// updates our Articles array to remove the
			// article
			fmt.Println("Event: " + Event.Id)
			EventList = append(EventList[:index], EventList[index+1:]...)
		}
	}
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	EventList = []Event{
		Event{Id: "1", Title: "Hello World", Place: "Malaga", Speaker: "Andrej", EventType: "Webinar", DateTime: "24.9.2018"},
		Event{Id: "2", Title: "Hello Death", Place: "Hell", Speaker: "Trump", EventType: "Live", DateTime: "6.6.1489"},
	}

	handleRequests()
}
