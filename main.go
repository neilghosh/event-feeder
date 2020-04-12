package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/neilghosh/event-feeder/model"
	"github.com/neilghosh/event-feeder/service"
	"github.com/neilghosh/event-feeder/testing"
)

func restHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
    case "GET":
		w.WriteHeader(http.StatusOK)
		ids, ok := r.URL.Query()["id"]
	    var id string

		if !ok || len(ids[0]) < 1 {
			log.Println("Url Param 'id' is missing")			
		} else {
			id = ids[0]
		}

		var response = service.GetFeed(id)
		json.NewEncoder(w).Encode(response)
        //w.Write([]byte(`{"message": "get called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		var request model.FeedItemRequest

		// This is another way to get the body 
		// body, er := ioutil.ReadAll(r.Body)
		// if er != nil {
		// 	panic(er)
		// }
		// log.Println(string(body))
		// err = json.Unmarshal(body, &request)

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		var response = service.PostFeed(request)
		json.NewEncoder(w).Encode(response)
    case "PUT":
		ids, ok := r.URL.Query()["id"]
	    var id string

		if !ok || len(ids[0]) < 1 {
			log.Println("Url Param 'id' is missing")			
		} else {
			id = ids[0]
		}

        w.WriteHeader(http.StatusAccepted)
		var request model.FeedItemRequest

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		var response = service.UpdateFeed(id, request)
		json.NewEncoder(w).Encode(response)
	case "DELETE":
		
		ids, ok := r.URL.Query()["id"]
	    var id string

		if !ok || len(ids[0]) < 1 {
			log.Println("Url Param 'id' is missing")			
		} else {
			id = ids[0]
		}
		service.DeleteFeed(id)
        w.WriteHeader(http.StatusNoContent)
        w.Write([]byte(`{"message": "delete called"}`))
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "not found"}`))
	}
	
}

func main() {
	http.HandleFunc("/api/", restHandler)
	http.HandleFunc("/test", testing.TestHandler)
	http.HandleFunc("/", testing.EchoHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
