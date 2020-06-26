package main

import (
	"encoding/json"
	"hello/config"
	"hello/db"
	"hello/models"
	"hello/view"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Pin  *view.Pin `json:"pin"`
	Tags []string  `json:"tags"`
}

type Response struct {
	Pin  *view.Pin   `json:"pin"`
	Tags []*view.Tag `json:"tags"`
}

func hello(event Event) (string, error) {
	c, err := config.ReadConfig()
	if err != nil {
		return "", err
	}

	db, err := db.ConnectToMySql(&c.DB)
	if err != nil {
		return "", err
	}

	data := models.NewSQLDataStorage(db)

	tags, err := data.GetTagsByTagNames(event.Tags)
	if err != nil {
		return "", err
	}

	now := time.Now()
	for _, tagName := range event.Tags {
		var existInDB bool

		for _, tag := range tags {
			if tagName == tag.Tag {
				existInDB = true
				break
			}
		}

		if !existInDB {
			newTag := &models.Tag{
				Tag:       tagName,
				CreatedAt: now,
				UpdatedAt: now,
			}

			newTag, _ = data.CreateTag(newTag)
			tags = append(tags, newTag)
		}
	}

	for _, tag := range tags {
		err = data.CreatePinTag(event.Pin.ID, tag.ID)
		if err != nil {
			return "", err
		}
	}

	response := Response{
		Pin:  event.Pin,
		Tags: view.NewTags(tags),
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func main() {
	lambda.Start(hello)
}
