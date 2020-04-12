package model

import (	
	"time"
)

type FeedItem struct {
	FeedName string
	FeedItemId string
	Active   bool
	Content  string
	Created  time.Time
	EventDate time.Time
}

type FeedItemRequest struct {
	Name string 
	Content string `json:"content"`
	EventDate time.Time
}

type FeedItemResponse struct {
	Id string 
}