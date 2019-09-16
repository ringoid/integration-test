package apimodel

import (
	"fmt"
)

const (
	Region     = "eu-west-1"
	MaxRetries = 3

	SecretWordKeyName = "secret-word-key"
	SecretWordKeyBase = "%s/SecretWord"

	InternalServerError           = `{"errorCode":"InternalServerError","errorMessage":"Internal Server Error"}`
	WrongRequestParamsClientError = `{"errorCode":"WrongParamsClientError","errorMessage":"Wrong request params"}`
	PhoneNumberClientError        = `{"errorCode":"PhoneNumberClientError","errorMessage":"Phone number is invalid"}`
	CountryCallingCodeClientError = `{"errorCode":"CountryCallingCodeClientError","errorMessage":"Country code is invalid"}`

	WrongSessionIdClientError        = `{"errorCode":"WrongSessionIdClientError","errorMessage":"Session id is invalid"}`
	NoPendingVerificationClientError = `{"errorCode":"NoPendingVerificationClientError","errorMessage":"No pending verifications found"}`
	WrongVerificationCodeClientError = `{"errorCode":"WrongVerificationCodeClientError","errorMessage":"Wrong verification code"}`

	WrongYearOfBirthClientError   = `{"errorCode":"WrongYearOfBirthClientError","errorMessage":"Wrong year of birth"}`
	WrongSexClientError           = `{"errorCode":"WrongSexClientError","errorMessage":"Wrong sex"}`
	InvalidAccessTokenClientError = `{"errorCode":"InvalidAccessTokenClientError","errorMessage":"Invalid access token"}`

	TooOldAppVersionClientError = `{"errorCode":"TooOldAppVersionClientError","errorMessage":"Too old app version"}`
)

type BaseResponse struct {
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (resp BaseResponse) String() string {
	return fmt.Sprintf("[BaseResponse={errorCode=%s, errorMessage=%s}", resp.ErrorCode, resp.ErrorMessage)
}

var PhotoResolution480x640 = "480x640"
var PhotoResolution720x960 = "720x960"
var PhotoResolution1080x1440 = "1080x1440"
var PhotoResolution1440x1920 = "1440x1920"

const (
	NewFacesSourceFeed   = "new_faces"
	WhoLikedMeSourceFeed = "who_liked_me"
	MatchesSourceFeed    = "matches"
	MessagesSourceFeed   = "messages"

	ViewActionType     = "VIEW"
	LikeActionType     = "LIKE"
	UnlikeActionType   = "UNLIKE"
	BlockActionType    = "BLOCK"
	MessageActionType  = "MESSAGE"
	LocationActionType = "LOCATION"
	ReadMessageActionType = "READ_MESSAGE"
)

type Bot struct {
	BotId          string
	BotAccessToken string
	IsPassive      bool
}

func (obj Bot) String() string {
	return fmt.Sprintf("%#v", obj)
}
