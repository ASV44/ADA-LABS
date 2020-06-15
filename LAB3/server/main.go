package main

import (
	"encoding/json"
	"fmt"
	"github.com/ASV44/ADA-LAB3/db"
	"github.com/ASV44/ADA-LAB3/server/models"
	"github.com/ASV44/ADA-LAB3/server/protobuf"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	database, err := db.NewDatabaseConnection("user:password@tcp(db:3306)/lab3")
	if err != nil {
		fmt.Println("Failed to connect to database", err)
		return
	}

	for i := 0; i < 1000; i++ {
		//event, err := getEventWithJsonPayload(i)
		event, err := getEventWithProtobufPayload(i)
		if err != nil {
			log.Fatalf("Failed to convert to generate data %s", err)
		}

		database.InsertEvent(event)
	}

	println("Successfully write all data")

	forever := make(chan bool)
	<-forever
}

func getEventWithJsonPayload(index int) (db.Event, error) {
	entityType, payload := getJsonPayload(index)
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

func getJsonPayload(index int) (string, interface{}) {
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

func getEventWithProtobufPayload(index int) (db.Event, error) {
	entityType, payload := getProtobufPayload(index)
	protoBufData, err := proto.Marshal(payload)
	if err != nil {
		return db.Event{}, err
	}

	return db.Event{
		ID:      index,
		Entity:  entityType,
		Payload: protoBufData,
	}, nil
}

func getProtobufPayload(index int) (string, proto.Message) {
	index32 := int32(index)
	switch {
	case index < 333:
		return db.Task, &protobuf.Task{
			Id:          index32,
			Version:     0,
			Title:       "Random title",
			Description: "Random description",
			StatusId:    "0bc915fc-be9e-411d-8077-ff5212e09c3a",
			UserId:      "0bc915fc-be9e-411d-8077-ff5212e09c3a",
			CreatedAt:   "2020-06-14T15:44",
		}
	case index >= 333 && index < 667:
		return db.TaskComment, &protobuf.TaskComment{
			Id:        index32,
			Version:   0,
			Content:   "Random Content",
			CreatedAt: "2020-06-14T15:44",
		}
	case index >= 667:
		return db.User, &protobuf.User{
			Id:   index32,
			Name: "Random Name",
		}
	default:
		return "", nil
	}
}
