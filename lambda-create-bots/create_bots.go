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
)

type CreateBotsRequest struct {
	Sex        string `json:"sex"`
	PassiveNum int    `json:"passiveNum"`
	ActiveNum  int    `json:"activeNum"`
}

func (obj CreateBotsRequest) String() string {
	return fmt.Sprintf("%#v", obj)
}

func handler(ctx context.Context, request CreateBotsRequest) (error) {
	lc, _ := lambdacontext.FromContext(ctx)
	return createBots(request.PassiveNum, request.ActiveNum, apimodel.BotsTableName, apimodel.AwsDynamoDB, lc)
}

func main() {
	basicLambda.Start(handler)
}

func createBots(passiveNum, activeNum int, tableName string, awsDb *dynamodb.DynamoDB, lc *lambdacontext.LambdaContext) error {
	apimodel.Anlogger.Debugf(lc, "create_bots.go : create bots passive [%d] and active [%d]", passiveNum, activeNum)

	for i := 0; i < passiveNum; i++ {
		userId, token := generateCatDogBot(false, "female", lc)
		bot := apimodel.Bot{
			BotId:          userId,
			BotAccessToken: token,
			IsPassive:      true,
		}
		writeItem(bot, tableName, awsDb, lc)
	}

	for i := 0; i < activeNum; i++ {
		userId, token := generateCatDogBot(true, "female", lc)
		bot := apimodel.Bot{
			BotId:          userId,
			BotAccessToken: token,
			IsPassive:      false,
		}
		writeItem(bot, tableName, awsDb, lc)
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

func generateCatDogBot(active bool, sex string, lc *lambdacontext.LambdaContext) (string, string) {
	token := apitests.CreateUser(sex, lc)
	userId := apimodel.UserId(token, lc)

	time.Sleep(time.Second)
	apitests.UploadCatDogImage(token, sex, active, lc)
	time.Sleep(1 * time.Second)
	apitests.UploadCatDogImage(token, sex, active, lc)
	time.Sleep(1 * time.Second)
	apitests.UploadCatDogImage(token, sex, active, lc)

	return userId, token
}
