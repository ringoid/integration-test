package apitests

import (
	"github.com/aws/aws-lambda-go/lambdacontext"
	"../apimodel"
)

func ActionsLocation(token string, lat, lon float64, lc *lambdacontext.LambdaContext) {
	actions := []apimodel.Action{apimodel.Action{
		ActionType: apimodel.LocationActionType,
		Lat:        lat,
		Lon:        lon,
		ActionTime: 123,
	}}
	resp := apimodel.Actions(token, actions, true, lc)
	if len(resp.ErrorCode) != 0 {
		panic("error when send location action")
	}
}

func ActionsWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	actions := []apimodel.Action{apimodel.Action{
		ActionType:     apimodel.LikeActionType,
		SourceFeed:     apimodel.NewFacesSourceFeed,
		TargetPhotoId:  "targetPhotoId",
		TargetUserId:   "targetUserId",
		LikeCount:      1,
		ViewCount:      1,
		ViewTimeMillis: 2,
		ActionTime:     123,
	}}
	resp := apimodel.Actions(token, actions, false, lc)
	if resp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /actions with old build num")
	}
}

func ActionsWithOldAccessToken(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	DeleteUser(token, lc)
	actions := []apimodel.Action{apimodel.Action{
		ActionType:     apimodel.LikeActionType,
		SourceFeed:     apimodel.NewFacesSourceFeed,
		TargetPhotoId:  "targetPhotoId",
		TargetUserId:   "targetUserId",
		LikeCount:      1,
		ViewCount:      1,
		ViewTimeMillis: 2,
		ActionTime:     123,
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
		ActionType:     apimodel.LikeActionType,
		TargetPhotoId:  "targetPhotoId",
		TargetUserId:   "targetUserId",
		LikeCount:      1,
		ViewCount:      1,
		ViewTimeMillis: 2,
		ActionTime:     123,
	}}
	resp := apimodel.Actions(token, actions, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /actions with empty source feed")
	}
}

func ActionsWithWrongSourceFeed(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	actions := []apimodel.Action{apimodel.Action{
		SourceFeed:     apimodel.NewFacesSourceFeed + "s",
		ActionType:     apimodel.LikeActionType,
		TargetPhotoId:  "targetPhotoId",
		TargetUserId:   "targetUserId",
		LikeCount:      1,
		ViewCount:      1,
		ViewTimeMillis: 2,
		ActionTime:     123,
	}}
	resp := apimodel.Actions(token, actions, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /actions with wrong source feed")
	}
}

func ActionsWithEmptyActionType(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	actions := []apimodel.Action{apimodel.Action{
		SourceFeed:     apimodel.NewFacesSourceFeed,
		TargetPhotoId:  "targetPhotoId",
		TargetUserId:   "targetUserId",
		LikeCount:      1,
		ViewCount:      1,
		ViewTimeMillis: 2,
		ActionTime:     123,
	}}
	resp := apimodel.Actions(token, actions, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /actions with empty action type")
	}
}

func ActionsWithWrongActionType(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	actions := []apimodel.Action{apimodel.Action{
		SourceFeed:     apimodel.NewFacesSourceFeed,
		ActionType:     apimodel.LikeActionType + "s",
		TargetPhotoId:  "targetPhotoId",
		TargetUserId:   "targetUserId",
		LikeCount:      1,
		ViewCount:      1,
		ViewTimeMillis: 2,
		ActionTime:     123,
	}}
	resp := apimodel.Actions(token, actions, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /actions with wrong action type")
	}
}

func ActionsWithEmptyTargetUserId(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	actions := []apimodel.Action{apimodel.Action{
		SourceFeed:     apimodel.NewFacesSourceFeed,
		ActionType:     apimodel.LikeActionType,
		TargetPhotoId:  "targetPhotoId",
		LikeCount:      1,
		ViewCount:      1,
		ViewTimeMillis: 2,
		ActionTime:     123,
	}}
	resp := apimodel.Actions(token, actions, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /actions with empty target user id")
	}
}

func ActionsWithEmptyTargetPhotoId(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	actions := []apimodel.Action{apimodel.Action{
		SourceFeed:     apimodel.NewFacesSourceFeed,
		ActionType:     apimodel.LikeActionType,
		TargetUserId:   "targetUserId",
		LikeCount:      1,
		ViewCount:      1,
		ViewTimeMillis: 2,
		ActionTime:     123,
	}}
	resp := apimodel.Actions(token, actions, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /actions with empty target photo id")
	}
}
