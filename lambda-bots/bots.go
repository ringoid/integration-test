package main

import (
	"context"
	basicLambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"../apimodel"
	"github.com/ringoid/commons"
	"encoding/json"
	"fmt"
	"time"
)

var bots []apimodel.Bot

var timeForBotUpdate int64

func handler(ctx context.Context, event events.SQSEvent) (error) {
	lc, _ := lambdacontext.FromContext(ctx)
	apimodel.Anlogger.Infof(lc, "bots.go : start bots function")
	var err error
	if timeForBotUpdate == 0 || time.Now().Unix()-timeForBotUpdate > 5 || len(bots) == 0 {
		timeForBotUpdate = time.Now().Unix()
		bots, err = getAllBot(apimodel.BotsTableName, apimodel.AwsDynamoDB, nil)
		if err != nil {
			return err
		}
	}

	for _, record := range event.Records {
		body := record.Body
		var aEvent commons.BaseInternalEvent
		err := json.Unmarshal([]byte(body), &aEvent)
		if err != nil {
			apimodel.Anlogger.Errorf(lc, "bots.go : error unmarshal body [%s] to BaseInternalEvent : %v", body, err)
			return fmt.Errorf("error unmarshal body %s : %v", body, err)
		}
		apimodel.Anlogger.Debugf(lc, "bots.go : handle record %v", aEvent)

		switch aEvent.EventType {
		case "BOT_ACTION_USER_LIKE_PHOTO":
			apimodel.Anlogger.Debugf(lc, "bots.go : handle event type %s", "BOT_ACTION_USER_LIKE_PHOTO")
			if err = like(body, bots, lc); err != nil {
				return err
			}
		case "BOT_ACTION_USER_VIEW_PHOTO":
			apimodel.Anlogger.Debugf(lc, "bots.go : handle event type %s", "BOT_ACTION_USER_VIEW_PHOTO")
		case "BOT_ACTION_USER_BLOCK_OTHER":
			apimodel.Anlogger.Debugf(lc, "bots.go : handle event type %s", "BOT_ACTION_USER_BLOCK_OTHER")
		case "BOT_ACTION_USER_UNLIKE_PHOTO":
			apimodel.Anlogger.Debugf(lc, "bots.go : handle event type %s", "BOT_ACTION_USER_UNLIKE_PHOTO")
		case "BOT_ACTION_USER_MESSAGE":
			apimodel.Anlogger.Debugf(lc, "bots.go : handle event type %s", "BOT_ACTION_USER_MESSAGE")
			if err = message(body, bots, lc); err != nil {
				return err
			}
		}
	}
	apimodel.Anlogger.Infof(lc, "bots.go : finish bots function")
	return nil
}

func main() {
	basicLambda.Start(handler)
}
