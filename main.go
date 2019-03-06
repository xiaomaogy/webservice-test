package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Features struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Url         string `json:"url"`
	ImageUrl    string `json:"image_url"`
}

type Action string

const (
	CLICK_ACTION Action = "click"
	VIEW_ACTION  Action = "view"
)

type InteractionRequest struct {
	Id     string `json:"id"`
	Action Action `json:"action"`
}

type Interaction struct {
	View  uint64 `json:"view"`
	Click uint64 `json:"click"`
}

type Ad struct {
	Features    Features           `json:"features"`
	Tags        map[string]float64 `json:"tags"`
	Interaction Interaction        `json:"interaction"`
}

func main() {
	fmt.Println("started-service")

	r := mux.NewRouter()
	r.HandleFunc("/ad", postHandler).Methods("POST")
	r.HandleFunc("/ad", listHandler).Methods("GET")
	r.HandleFunc("/ad/{id}", getHandler).Methods("GET")
	r.HandleFunc("/ad/interaction", interactionHandler).Methods("POST")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is post handler.")

	decoder := json.NewDecoder(r.Body)
	var f Features
	if err := decoder.Decode(&f); err != nil {
		http.Error(w, "Failed to parse JSON input", http.StatusBadRequest)
		fmt.Printf("Failed to parse JSON input %v\n", err)
		return
	}
	fmt.Fprintf(w, "Post received: %s\n", f.Description)
}

//func postHandler(w http.ResponseWriter, r *http.Request){}

//func handlerSearch(w http.ResponseWriter, r *http.Request){}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is get handler.")

	id := mux.Vars(r)["id"]
	fmt.Printf("Ad id is: %s", id)

	ad := &Ad{
		Features: Features{
			Id:          id,
			Description: "this is the second ad of laioffer",
			Url:         "www.laioffer.com",
			ImageUrl:    "www.laioffer.com/Images/2222",
		},
		Tags: map[string]float64{"foo": 1, "bar": 2},
		Interaction: Interaction{
			View:  1,
			Click: 0,
		},
	}

	js, err := json.Marshal(ad)
	if err != nil {
		m := fmt.Sprintf("Failed to parse ad object to JSON %v", err)
		http.Error(w, m, http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func interactionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is interaction handler.")

	decoder := json.NewDecoder(r.Body)
	var req InteractionRequest
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Failed to parse JSON input", http.StatusBadRequest)
		fmt.Printf("Failed to parse JSON input %v\n", err)
		return
	}
	fmt.Fprintf(w, "interaction received: %s\n", req.Action)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is list handler")

	//Return a list of fake ads
	var ads []*Ad
	ads = append(ads, &Ad{
		Features: Features{
			Id:          "1111",
			Description: "this is the first ad of laioffer",
			Url:         "www.laioffer.com",
			ImageUrl:    "www.laioffer.com/Images/1111",
		},
		Tags: map[string]float64{"foo": 1, "bar": 2},
		Interaction: Interaction{
			View:  2,
			Click: 1,
		},
	})

	ads = append(ads, &Ad{
		Features: Features{
			Id:          "2222",
			Description: "this is the second ad of laioffer",
			Url:         "www.laioffer.com",
			ImageUrl:    "www.laioffer.com/Images/2222",
		},
		Tags: map[string]float64{"foo": 1, "bar": 2},
		Interaction: Interaction{
			View:  1,
			Click: 0,
		},
	})

	js, err := json.Marshal(ads)
	if err != nil {
		m := fmt.Sprintf("Failed to parse ad objects to JSON %v", err)
		http.Error(w, m, http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
