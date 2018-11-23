package apitests

import (
	"github.com/aws/aws-lambda-go/lambdacontext"
	"../apimodel"
)

const (
	newFacesSourceFeed   = "new_faces"
	whoLikedMeSourceFeed = "who_liked_me"
	matchesSourceFeed    = "matches"
	messagesSourceFeed   = "messages"

	viewActionType   = "VIEW"
	likeActionType   = "LIKE"
	unlikeActionType = "UNLIKE"
	blockActionType  = "BLOCK"
)

func ActionsWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	actions := []apimodel.Action{apimodel.Action{
		ActionType:    likeActionType,
		SourceFeed:    newFacesSourceFeed,
		TargetPhotoId: "targetPhotoId",
		TargetUserId:  "targetUserId",
		LikeCount:     1,
		ViewCount:     1,
		ViewTimeSec:   2,
		ActionTime:    123,
	}}
	resp := apimodel.Actions(token, actions, false, lc)
	if resp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /actions with old build num")
	}
}

func ActionsWithOldAccessToken(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	apimodel.Logout(token, true, lc)
	actions := []apimodel.Action{apimodel.Action{
		ActionType:    likeActionType,
		SourceFeed:    newFacesSourceFeed,
		TargetPhotoId: "targetPhotoId",
		TargetUserId:  "targetUserId",
		LikeCount:     1,
		ViewCount:     1,
		ViewTimeSec:   2,
		ActionTime:    123,
	}}
	resp := apimodel.Actions(token, actions, true, lc)
	if resp.ErrorCode != "InvalidAccessTokenClientError" {
		panic("there is no InvalidAccessTokenClientError when call /actions with old access token")
	}
}

func ActionsWithNilActions(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	resp := apimodel.Actions(token, nil, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /actions with nil actions")
	}
}

func ActionsWithEmptySourceFeed(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	actions := []apimodel.Action{apimodel.Action{
		ActionType:    likeActionType,
		TargetPhotoId: "targetPhotoId",
		TargetUserId:  "targetUserId",
		LikeCount:     1,
		ViewCount:     1,
		ViewTimeSec:   2,
		ActionTime:    123,
	}}
	resp := apimodel.Actions(token, actions, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /actions with empty source feed")
	}
}

func ActionsWithWrongSourceFeed(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	actions := []apimodel.Action{apimodel.Action{
		SourceFeed:    newFacesSourceFeed + "s",
		ActionType:    likeActionType,
		TargetPhotoId: "targetPhotoId",
		TargetUserId:  "targetUserId",
		LikeCount:     1,
		ViewCount:     1,
		ViewTimeSec:   2,
		ActionTime:    123,
	}}
	resp := apimodel.Actions(token, actions, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /actions with wrong source feed")
	}
}

func ActionsWithEmptyActionType(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	actions := []apimodel.Action{apimodel.Action{
		SourceFeed:    newFacesSourceFeed,
		TargetPhotoId: "targetPhotoId",
		TargetUserId:  "targetUserId",
		LikeCount:     1,
		ViewCount:     1,
		ViewTimeSec:   2,
		ActionTime:    123,
	}}
	resp := apimodel.Actions(token, actions, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /actions with empty action type")
	}
}

func ActionsWithWrongActionType(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	actions := []apimodel.Action{apimodel.Action{
		SourceFeed:    newFacesSourceFeed,
		ActionType:    likeActionType + "s",
		TargetPhotoId: "targetPhotoId",
		TargetUserId:  "targetUserId",
		LikeCount:     1,
		ViewCount:     1,
		ViewTimeSec:   2,
		ActionTime:    123,
	}}
	resp := apimodel.Actions(token, actions, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /actions with wrong action type")
	}
}

func ActionsWithEmptyTargetUserId(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	actions := []apimodel.Action{apimodel.Action{
		SourceFeed:    newFacesSourceFeed,
		ActionType:    likeActionType,
		TargetPhotoId: "targetPhotoId",
		LikeCount:     1,
		ViewCount:     1,
		ViewTimeSec:   2,
		ActionTime:    123,
	}}
	resp := apimodel.Actions(token, actions, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /actions with empty target user id")
	}
}

func ActionsWithEmptyTargetPhotoId(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	actions := []apimodel.Action{apimodel.Action{
		SourceFeed:   newFacesSourceFeed,
		ActionType:   likeActionType,
		TargetUserId: "targetUserId",
		LikeCount:    1,
		ViewCount:    1,
		ViewTimeSec:  2,
		ActionTime:   123,
	}}
	resp := apimodel.Actions(token, actions, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /actions with empty target photo id")
	}
}
