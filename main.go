package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

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
var Articles []Article
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
	muxRouter.HandleFunc("/articles", returnAllArticles)
	muxRouter.HandleFunc("/getAllEvents", returnAllEvents)
	muxRouter.HandleFunc("/getEvent/{Id}", returnSingleEvent)
	log.Fatal(http.ListenAndServe(":10000", muxRouter))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
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

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	EventList = []Event{
		Event{Id: "1", Title: "Hello World", Place: "Malaga", Speaker: "Andrej", EventType: "Webinar", DateTime: "24.9.2018"},
		Event{Id: "2", Title: "Hello Death", Place: "Hell", Speaker: "Trump", EventType: "Live", DateTime: "6.6.1489"},
	}

	handleRequests()
}
