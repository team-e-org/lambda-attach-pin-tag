package main

import (
	"app/config"
	"app/db"
	"app/logs"
	"app/models"
	"app/view"
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"regexp"
	"time"

	l "github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Pin  *view.Pin `json:"pin"`
	Tags []string  `json:"tags"`
}

type Response struct {
	Pin  *view.Pin   `json:"pin"`
	Tags []*view.Tag `json:"tags"`
}

func handler(event Event) (string, error) {
	logs.Info("Lambda started, event: %v", event)

	c, err := config.ReadConfig()
	if err != nil {
		logs.Error("read db config: %v", err)
		return "", err
	}

	db, err := db.ConnectToMySql(&c.DB)
	if err != nil {
		logs.Error("connecting to mysql: %v", err)
		return "", err
	}

	data := models.NewSQLDataStorage(db)

	tags, err := data.GetTagsByTagNames(event.Tags)
	if err != nil {
		logs.Error("getting tags by tag names: %v", err)
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
			logs.Error("creating pin_tags: %v", err)
			return "", err
		}
	}

	event.Pin.ImageURL = removeURLDomain(event.Pin.ImageURL)
	response := Response{
		Pin:  event.Pin,
		Tags: view.NewTags(tags),
	}

	b, err := json.Marshal(response)
	if err != nil {
		logs.Error("serializing json: %v", err)
		return "", err
	}

	url := os.Getenv("URL")

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewReader(b),
	)
	if err != nil {
		logs.Error("Can not create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("Can not do request: %v", err)
	}
	defer resp.Body.Close()

	logs.Info("response: %v", resp)

	// svc := lambda.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	// input := &lambda.InvokeInput{
	// 	FunctionName:   aws.String("arn:aws:lambda:ap-northeast-1:444207867088:function:insertDynamo"),
	// 	Payload:        bytes,
	// 	InvocationType: aws.String("Event"),
	// }

	// resp, err := svc.Invoke(input)

	// if err != nil {
	// 	logs.Error("calling next lambda: %v", err)
	// 	return "", err
	// }
	// logs.Info("response from next lambda: %v", resp)

	return string(b), nil
}

func main() {
	l.Start(handler)
}

func removeURLDomain(url string) string {
	reg := regexp.MustCompile(`pins(.+)`)
	return reg.FindString(url)
}
