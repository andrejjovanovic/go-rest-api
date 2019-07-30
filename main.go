package main

import (
    "fmt"
    "log"
	"net/http"
	"encoding/json"
)

type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type Event struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Place string `json:"place"`
	Speaker string `json:"speaker"`
	EventType string `json:"eventType"`
	DateTime string `json:"dateTime"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article
var EventList []Event


func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	// add our articles route and map it to our 
    // returnAllArticles function like so
	http.HandleFunc("/articles", returnAllArticles)
	http.HandleFunc("/getAllEvents", returnAllEvents)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func returnAllEvents(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(EventList)
}


func main() {
    Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	
	EventList = []Event{
		Event{Id: "1", Title: "Hello World", Place: "Malaga", Speaker: "Andrej", EventType: "Webinar", DateTime: "24.9.2018"},
	}

    handleRequests()
}