package service

import(
	"context"
	"log"
	"fmt"
	"encoding/json"
	"time"

	"cloud.google.com/go/datastore"

	"github.com/google/uuid"

	"github.com/neilghosh/event-feeder/db"
	"github.com/neilghosh/event-feeder/model"
)

func DeleteFeed(id string) {
	ctx := context.Background()
	datastoreClient :=  db.GetDataStoreClient(ctx);

	dskey := datastore.NameKey("FeedItem", id, nil)
	if err := datastoreClient.Delete(ctx, dskey); err != nil {
		log.Printf(fmt.Sprintf("deleted entry to database %v", dskey))
	}
}

func GetFeed(id string) model.FeedItem {
	ctx := context.Background()
	datastoreClient :=  db.GetDataStoreClient(ctx);

	dskey := datastore.NameKey("FeedItem", id, nil)
	feedItem := model.FeedItem{}
	if err := datastoreClient.Get(ctx, dskey, &feedItem); err != nil {
		log.Printf(fmt.Sprintf("Written entry to database %v", dskey))
	}

	log.Printf(fmt.Sprintf("Written entry to database %v", dskey))
	return feedItem
} 

func  PostFeed(request model.FeedItemRequest) model.FeedItemResponse {
	log.Printf("Request: %+v", request)

	feedItem := &model.FeedItem{
		FeedName: request.Name,
		Active:  true,
		Content : request.Content, 
		Created: time.Now(),
		EventDate : request.EventDate,
	}

	key := uuid.New().String()		
	db.WriteToDatabase(key, feedItem)

	//taskStr, _ := json.Marshal(feedItem)

	response := model.FeedItemResponse{
		Id:     key,
	}

	data, _ := json.Marshal(response)
	log.Println("Respionse : "+string(data))
	return response
}

func  UpdateFeed(id string, request model.FeedItemRequest) model.FeedItemResponse {
	log.Printf("Request: %+v", request)

	feedItem := &model.FeedItem{
		FeedName: request.Name,
		Active:  true,
		Content : request.Content, 
		Created: time.Now(),
		EventDate : request.EventDate,
	}

	key := id		
	db.WriteToDatabase(key, feedItem)

	//taskStr, _ := json.Marshal(feedItem)

	response := model.FeedItemResponse{
		Id:     key,
	}

	data, _ := json.Marshal(response)
	log.Println("Respionse : "+string(data))
	return response
}
