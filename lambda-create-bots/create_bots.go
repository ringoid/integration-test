package main

import (
	"context"
	basicLambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"../apimodel"
	"../apitests"
	"time"
	"github.com/ringoid/commons"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"math/rand"
)

type CreateBotsRequest struct {
	Sex        string `json:"sex"`
	PassiveNum int    `json:"passiveNum"`
	ActiveNum  int    `json:"activeNum"`
	StartFrom  int    `json:"startNum"`
}

func (obj CreateBotsRequest) String() string {
	return fmt.Sprintf("%#v", obj)
}

func handler(ctx context.Context, request CreateBotsRequest) (error) {
	lc, _ := lambdacontext.FromContext(ctx)
	return createBots(request.StartFrom, request.PassiveNum, request.ActiveNum, request.Sex, apimodel.BotsTableName, apimodel.AwsDynamoDB, lc)
}

func main() {
	basicLambda.Start(handler)
}

func createBots(startNum, passiveNum, activeNum int, targetSex, tableName string, awsDb *dynamodb.DynamoDB, lc *lambdacontext.LambdaContext) error {
	apimodel.Anlogger.Debugf(lc, "create_bots.go : create bots passive [%d] and active [%d]", passiveNum, activeNum)

	photoCounter := 2
	for i := startNum; i < startNum+passiveNum; i++ {
		photoCounter += 2
		//userId, token := generateBot(false, targetSex, i, lc)
		userId, token := generateCatDogBot(false, targetSex, photoCounter, lc)
		updateBotsProfile(token, true, lc)
		apitests.ActionsLocation(token, 37.34299850463867, -122.05217742919922, lc)

		bot := apimodel.Bot{
			BotId:          userId,
			BotAccessToken: token,
			IsPassive:      true,
		}
		writeItem(bot, tableName, awsDb, lc)
		if photoCounter > 16 {
			photoCounter = 2
		}
	}

	photoCounter = 2
	for i := 0; i < activeNum; i++ {
		photoCounter += 2
		//userId, token := generateBot(true, targetSex, i, lc)
		userId, token := generateCatDogBot(true, targetSex, photoCounter, lc)
		updateBotsProfile(token, false, lc)
		apitests.ActionsLocation(token, 37.34299850463867, -122.05217742919922, lc)

		bot := apimodel.Bot{
			BotId:          userId,
			BotAccessToken: token,
			IsPassive:      false,
		}
		writeItem(bot, tableName, awsDb, lc)
		if photoCounter > 16 {
			photoCounter = 2
		}
	}

	apimodel.Anlogger.Debugf(lc, "create_bots.go : successfully create bots passive [%d] and active [%d]", passiveNum, activeNum)
	return nil
}

func writeItem(bot apimodel.Bot, tableName string, awsDb *dynamodb.DynamoDB, lc *lambdacontext.LambdaContext) error {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			commons.UserIdColumnName: {
				S: aws.String(bot.BotId),
			},
			"access_token": {
				S: aws.String(bot.BotAccessToken),
			},
			"passive": {
				BOOL: aws.Bool(bot.IsPassive),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := awsDb.PutItem(input)
	if err != nil {
		apimodel.Anlogger.Errorf(lc, "create_bots.go : error creating a bot %v in table [%s] : %v", bot, tableName, err)
		return fmt.Errorf("error creating a bots : %v", err)
	}
	return nil
}

func generateBot(active bool, sex string, baseNum int, lc *lambdacontext.LambdaContext) (string, string) {
	token := apitests.CreateUser(sex, lc)
	userId := apimodel.UserId(token, lc)

	time.Sleep(time.Second)
	prefix := fmt.Sprintf("Passive.%d", baseNum)
	if active {
		prefix = fmt.Sprintf("Active.%d", baseNum)
	}
	apitests.UploadImageWithPrefix(token, sex, prefix, 1, lc)
	time.Sleep(1 * time.Second)
	apitests.UploadImageWithPrefix(token, sex, prefix, 2, lc)
	time.Sleep(1 * time.Second)
	apitests.UploadImageWithPrefix(token, sex, prefix, 3, lc)

	return userId, token
}

func generateCatDogBot(active bool, sex string, photoNum int, lc *lambdacontext.LambdaContext) (string, string) {
	token := apitests.CreateUser(sex, lc)
	userId := apimodel.UserId(token, lc)

	for i := 0; i < photoNum; i++ {
		time.Sleep(time.Second)
		apitests.UploadCatDogImage(token, sex, active, lc)
	}
	return userId, token
}

func updateBotsProfile(token string, isItCat bool, lc *lambdacontext.LambdaContext) {
	property := rand.Intn(8) * 10
	transport := rand.Intn(8) * 10
	income := rand.Intn(5) * 10
	height := 150 + rand.Intn(50)
	educationLevel := rand.Intn(7) * 10
	hairColor := rand.Intn(7) * 10
	children := rand.Intn(5) * 10

	apitests.UpdateProfile(token, isItCat, property, transport, income, height, educationLevel, hairColor, children, lc)
}
