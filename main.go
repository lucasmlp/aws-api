package main

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

var (
	region = "us-west-2"
)

func main() {
	config := aws.Config{
		Region: &region,
	}

	session, err := session.NewSession(&config)
	if err != nil {
		log.Fatalln(err)
	}

	svc := apigateway.New(session, &config)

	apis, err := svc.GetRestApis(nil)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("apis: %v\n", apis)

	stageName := "dev"

	for _, api := range apis.Items {
		log.Printf("api: %v\n", api)
		getStageInput := apigateway.GetStageInput{
			RestApiId: api.Id,
			StageName: &stageName,
		}
		stage, err := svc.GetStage(&getStageInput)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("stage: %v\n", stage)
	}

	time.Sleep(10 * time.Second)
}
