package main

import (
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"../apimodel"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"github.com/ringoid/commons"
	"time"
	"math/rand"
	"github.com/satori/go.uuid"
)

func getAllBot(tableName string, awsDb *dynamodb.DynamoDB, lc *lambdacontext.LambdaContext) ([]apimodel.Bot, error) {
	apimodel.Anlogger.Debugf(lc, "bots_common.go : read all bots from the table")
	input := &dynamodb.ScanInput{
		ConsistentRead: aws.Bool(true),
		TableName:      aws.String(tableName),
	}

	result, err := awsDb.Scan(input)
	if err != nil {
		apimodel.Anlogger.Errorf(lc, "bots_common.go : error scan table [%s] : %v", tableName, err)
		return nil, fmt.Errorf("error scan table : %v", err)
	}

	bots := make([]apimodel.Bot, 0)
	for _, eachItem := range result.Items {
		bots = append(bots, apimodel.Bot{
			BotId:          *eachItem[commons.UserIdColumnName].S,
			BotAccessToken: *eachItem["access_token"].S,
			IsPassive:      *eachItem["passive"].BOOL,
		})
	}
	apimodel.Anlogger.Debugf(lc, "bots_common.go : successfully read all bots [%d] from the table", len(bots))
	return bots, nil
}

func wakeUpActiveBots(all []apimodel.Bot, lc *lambdacontext.LambdaContext) error {
	apimodel.Anlogger.Debugf(lc, "bots_common.go : wakeUp active bots")
	for _, bot := range all {
		if !bot.IsPassive {
			apimodel.Anlogger.Debugf(lc, "bots_common.go : found active bot %v", bot)

			newFaces := apimodel.GetNewFaces(bot.BotAccessToken, apimodel.PhotoResolution1440x1920, 0, true, lc)
			for _, each := range newFaces.Profiles {
				apimodel.Anlogger.Debugf(lc, "bots_common.go : send a like from newFaces for %v", each)

				//send like to random photo
				randPhotoIndex := 0
				if len(each.Photos) > 1 {
					randPhotoIndex = rand.Intn(len(each.Photos) - 1)
				}
				targetPhoto := each.Photos[randPhotoIndex]

				actions := []apimodel.Action{
					apimodel.Action{
						SourceFeed:     apimodel.NewFacesSourceFeed,
						ActionType:     apimodel.ViewActionType,
						TargetPhotoId:  targetPhoto.PhotoId,
						TargetUserId:   each.UserId,
						LikeCount:      0,
						ViewCount:      1,
						ViewTimeMillis: 2,
						ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
					},
					apimodel.Action{
						SourceFeed:    apimodel.NewFacesSourceFeed,
						ActionType:    apimodel.LikeActionType,
						TargetPhotoId: targetPhoto.PhotoId,
						TargetUserId:  each.UserId,
						LikeCount:     1,
						ActionTime:    time.Now().Round(time.Millisecond).UnixNano() / 1000000,
					},
				}
				apimodel.Actions(bot.BotAccessToken, actions, true, lc)
			}

			gLc := apimodel.GetLc(bot.BotAccessToken, apimodel.PhotoResolution1440x1920, 0, true, lc)
			for _, each := range gLc.LikesYou {
				apimodel.Anlogger.Debugf(lc, "bots_common.go : send a like from likesYou for %v", each)

				//send like to random photo
				randPhotoIndex := 0
				if len(each.Photos) > 1 {
					randPhotoIndex = rand.Intn(len(each.Photos) - 1)
				}
				targetPhoto := each.Photos[randPhotoIndex]

				actions := []apimodel.Action{
					apimodel.Action{
						SourceFeed:     apimodel.WhoLikedMeSourceFeed,
						ActionType:     apimodel.ViewActionType,
						TargetPhotoId:  targetPhoto.PhotoId,
						TargetUserId:   each.UserId,
						LikeCount:      0,
						ViewCount:      1,
						ViewTimeMillis: 2,
						ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
					},
					apimodel.Action{
						SourceFeed:    apimodel.WhoLikedMeSourceFeed,
						ActionType:    apimodel.LikeActionType,
						TargetPhotoId: targetPhoto.PhotoId,
						TargetUserId:  each.UserId,
						LikeCount:     1,
						ActionTime:    time.Now().Round(time.Millisecond).UnixNano() / 1000000,
					},
				}
				apimodel.Actions(bot.BotAccessToken, actions, true, lc)
			}
			for _, each := range gLc.Messages {
				apimodel.Anlogger.Debugf(lc, "bots_common.go : send a message from matches for %v", each)

				//send like to random photo
				randPhotoIndex := 0
				if len(each.Photos) > 1 {
					randPhotoIndex = rand.Intn(len(each.Photos) - 1)
				}
				targetPhoto := each.Photos[randPhotoIndex]
				clientMsgId, _ := uuid.NewV4()
				textFromBot := fmt.Sprintf("Bot [%s] at [%v] with clientMsgId [%s]", bot.BotId[0:4],
					time.Now().Format("15:04:05.000"), clientMsgId.String()[0:4])
				//randomText := fmt.Sprintf("Message from a bot (rand num %d)", rand.Intn(100))

				//generate read for each message
				readMsgs := make([]apimodel.Action, 0)
				for _, msg := range each.Messages {
					if !msg.WasYouSender && !msg.HaveBeenRead {
						readMsgs = append(readMsgs,
							apimodel.Action{
								SourceFeed:     apimodel.MessagesSourceFeed,
								ActionType:     apimodel.ReadMessageActionType,
								OppositeUserId: each.UserId,
								MessageId:      msg.MessageId,
								ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000})
					}
				}

				actions := []apimodel.Action{
					apimodel.Action{
						SourceFeed:     apimodel.MessagesSourceFeed,
						ActionType:     apimodel.ViewActionType,
						TargetPhotoId:  targetPhoto.PhotoId,
						TargetUserId:   each.UserId,
						LikeCount:      0,
						ViewCount:      1,
						ViewTimeMillis: 2,
						ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
					},
					apimodel.Action{
						SourceFeed:      apimodel.MessagesSourceFeed,
						ActionType:      apimodel.MessageActionType,
						TargetPhotoId:   targetPhoto.PhotoId,
						TargetUserId:    each.UserId,
						Text:            textFromBot,
						ClientMessageId: clientMsgId.String(),
						ActionTime:      time.Now().Round(time.Millisecond).UnixNano() / 1000000,
					},
				}
				actions = append(actions, readMsgs...)
				apimodel.Actions(bot.BotAccessToken, actions, true, lc)
			}

		}
	}
	apimodel.Anlogger.Debugf(lc, "bots_common.go : successfully finish active bots jobs")
	return nil
}
