package main

import (
	"github.com/aws/aws-lambda-go/lambdacontext"
	"../apimodel"
	"github.com/ringoid/commons"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
	"github.com/satori/go.uuid"
)

func message(body string, bots []apimodel.Bot, lc *lambdacontext.LambdaContext) error {
	apimodel.Anlogger.Debugf(lc, "message.go : start handle message event %v", body)
	var aEvent commons.UserMsgEvent
	err := json.Unmarshal([]byte(body), &aEvent)
	if err != nil {
		apimodel.Anlogger.Errorf(lc, "message.go : error unmarshal body [%s] to UserMsgEvent : %v", body, err)
		return fmt.Errorf("error unmarshal body %s : %v", body, err)
	}

	//first check if source is a bot
	for _, each := range bots {
		if each.BotId == aEvent.UserId {
			apimodel.Anlogger.Debugf(lc, "message.go : receive like from a bot, so return")
			return nil
		}
	}

	//second check who was target like
	var targetBot apimodel.Bot
	for _, each := range bots {
		if each.BotId == aEvent.TargetUserId {
			targetBot = each
			break
		}
	}

	if len(targetBot.BotId) == 0 {
		apimodel.Anlogger.Debugf(lc, "message.go : successfully finish handle message event %v, bot was not found", body)
		return nil
	}

	//ok we found the bot
	apimodel.Anlogger.Debugf(lc, "message.go : found a bot which receive message %v", targetBot)

	resp := apimodel.GetLc(targetBot.BotAccessToken, apimodel.PhotoResolution1440x1920, 0, true, lc)

	sourceFeed := "no"
	var profile apimodel.Profile
	for _, each := range resp.LikesYou {
		if each.UserId == aEvent.UserId {
			sourceFeed = apimodel.WhoLikedMeSourceFeed
			profile = each
		}
	}

	messages := make([]apimodel.Message, 0)
	if sourceFeed == "no" {
		for _, each := range resp.Messages {
			if each.UserId == aEvent.UserId {
				sourceFeed = apimodel.MessagesSourceFeed
				profile = each
				messages = each.Messages
			}
		}
	}

	if sourceFeed == "no" {
		apimodel.Anlogger.Errorf(lc, "message.go : error, can not find source user id [%s] in llm response %v", aEvent.UserId, resp)
		return fmt.Errorf("can not find user id [%s] in llm response", aEvent.UserId)
	}

	//send like to random photo
	randPhotoIndex := 0
	if len(profile.Photos) > 1 {
		randPhotoIndex = rand.Intn(len(profile.Photos) - 1)
	}
	targetPhoto := profile.Photos[randPhotoIndex]

	clientMsgId, _ := uuid.NewV4()
	textFromBot := fmt.Sprintf("Bot [%s] at [%v] replying to [%s] with clientMsgId [%s]",
		targetBot.BotId[0:4], time.Now().Format("15:04:05.000"), aEvent.Text, clientMsgId.String()[0:4])
	//randomText := fmt.Sprintf("Message from a bot (rand num %d)", rand.Intn(100))
	actions := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     sourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  targetPhoto.PhotoId,
			TargetUserId:   aEvent.UserId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:      sourceFeed,
			ActionType:      apimodel.MessageActionType,
			TargetPhotoId:   targetPhoto.PhotoId,
			TargetUserId:    aEvent.UserId,
			Text:            textFromBot,
			ClientMessageId: clientMsgId.String(),
			ActionTime:      time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}

	//generate read for each message
	readMsgs := make([]apimodel.Action, 0)
	for _, msg := range messages {
		if !msg.WasYouSender && !msg.HaveBeenRead {
			readMsgs = append(readMsgs,
				apimodel.Action{
					SourceFeed:     apimodel.MessagesSourceFeed,
					ActionType:     apimodel.ReadMessageActionType,
					OppositeUserId: aEvent.UserId,
					MessageId:      msg.MessageId,
					ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000})
		}
	}

	actions = append(actions, readMsgs...)
	apimodel.Actions(targetBot.BotAccessToken, actions, true, lc)

	wakeUpActiveBots(bots, lc)
	return nil
}
