package testing

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/neilghosh/event-feeder/constants"
)

type Test struct {
	Data string
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello World")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	//https://github.com/GoogleCloudPlatform/golang-samples/blob/master/datastore/snippets/snippet_test.go
	ctx := context.Background()
	log.Printf(fmt.Sprintf("Request Context %v", ctx))

	datastoreClient, err := datastore.NewClient(ctx, constants.PROJECT_ID)
	if err != nil {
		log.Printf(fmt.Sprintf("Datatore client error %v", err))	
		return
	} 
	
	log.Printf(fmt.Sprintf("Datatore client %v", datastoreClient))	
	

	dskey := datastore.NameKey("Test", "testKey", nil)
	log.Printf(fmt.Sprintf("Datatore Key %v", dskey))	


	testData := &Test{
		Data: "HelloWorld",
	}

	dskey, err = datastoreClient.Put(ctx, dskey, testData)
	if err != nil {
		log.Printf(fmt.Sprintf("Datatore writting error %v", err))	
		return
	} 
	log.Printf(fmt.Sprintf("Datastore Response %v", dskey))	

	fmt.Fprint(w, fmt.Sprintf("Datastore Response %v", dskey))
}
