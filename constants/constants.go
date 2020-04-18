package constants

import(
	"os"
	"log"
)

//const PROJECT_ID = "demoneil"
var projectId = ""

func init() {
	// Init ConfigMap here
	//get from env
	projectId = os.Getenv("GOOGLE_CLOUD_PROJECT");
	if projectId == "" {
		log.Println("GOOGLE_CLOUD_PROJECT is not set , so will try to use local emulator at " + os.Getenv("DATASTORE_EMULATOR_HOST"))			
		projectId = "project-id"
	}

}

func GetProject() string {
	return projectId
} 

const FEED_ENTITY = "FeedItem"