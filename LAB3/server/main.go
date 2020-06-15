package main

import (
	"encoding/json"
	"fmt"
	"github.com/ASV44/ADA-LAB3/db"
	"github.com/ASV44/ADA-LAB3/server/models"
	"log"
)

func main() {
	database, err := db.NewDatabaseConnection("user:password@tcp(db:3306)/lab3")
	if err != nil {
		fmt.Println("Failed to connect to database", err)
		return
	}

	for i := 0; i < 1000; i++ {
		event, err := getEventWithJsonPayload(i)
		if err != nil {
			log.Fatalf("Failed to convert to json %s", err)
		}

		database.InsertEvent(event)
	}

	println("Successfully write all data")

	forever := make(chan bool)
	<-forever
}

func getEventWithJsonPayload(index int) (db.Event, error) {
	entityType, payload := getEventJsonPayload(index)
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return db.Event{}, err
	}

	return db.Event{
		ID:      index,
		Entity:  entityType,
		Payload: jsonData,
	}, nil
}

func getEventJsonPayload(index int) (string, interface{}) {
	switch {
	case index < 333:
		return db.Task, models.Task{
			Id:          index,
			Version:     0,
			Title:       "Random title",
			Description: "Random description",
			StatusID:    "0bc915fc-be9e-411d-8077-ff5212e09c3a",
			UserID:      "0bc915fc-be9e-411d-8077-ff5212e09c3a",
			CreatedAt:   "2020-06-14T15:44",
		}
	case index >= 333 && index < 667:
		return db.TaskComment, models.TaskComment{
			Id:        index,
			Version:   0,
			Content:   "Random Content",
			CreatedAt: "2020-06-14T15:44",
		}
	case index >= 667:
		return db.User, models.User{
			Id:   index,
			Name: "Random Name",
		}
	default:
		return "", nil
	}
}
