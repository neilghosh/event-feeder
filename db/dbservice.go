package db

import (
	"fmt"
	"context"
	"log"

	"cloud.google.com/go/datastore"

	"github.com/neilghosh/event-feeder/model"
	"github.com/neilghosh/event-feeder/constants"
)

func GetDataStoreClient(ctx context.Context) *datastore.Client {
	// Creates a datastore client.
	datastoreClient, err := datastore.NewClient(ctx, constants.PROJECT_ID)
	if err != nil {
		log.Fatal(err)	
	}
	return datastoreClient

}

func WriteToDatabase(key string,feedItem *model.FeedItem){
	ctx := context.Background()
	datastoreClient := GetDataStoreClient(ctx)
	dskey := datastore.NameKey(constants.FEED_ENTITY, key, nil)
	dskey, _ = datastoreClient.Put(ctx, dskey, feedItem)
	log.Printf(fmt.Sprintf("Written entry to database %v", dskey))
}