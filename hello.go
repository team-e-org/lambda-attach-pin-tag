package main

import (
	"encoding/json"
	"fmt"
	"hello/config"
	"hello/db"
	"hello/models"
	"hello/view"
	"time"

	l "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
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

	svc := lambda.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	input := &lambda.InvokeInput{
		FunctionName:   aws.String("arn:aws:lambda:ap-northeast-1:444207867088:function:insertDynamo"),
		Payload:        bytes,
		InvocationType: aws.String("Event"),
	}

	resp, err := svc.Invoke(input)
	if err != nil {
		return "", err
	}
	fmt.Println(resp)

	return string(bytes), nil
}

func main() {
	l.Start(hello)
}
