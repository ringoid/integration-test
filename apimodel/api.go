package apimodel

import (
	"fmt"
)

type CreateUsersRequest struct {
	NextNum  int `json:"nextNum"`
	MenNum   int `json:"menNum"`
	WomenNum int `json:"womenNum"`
}

func (req CreateUsersRequest) String() string {
	return fmt.Sprintf("%#v", req)
}

//Request - Response model
type AuthResp struct {
	BaseResponse
	SessionId  string `json:"sessionId"`
	CustomerId string `json:"customerId"`
}

func (resp AuthResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

type StartReq struct {
	WarmUpRequest              bool   `json:"warmUpRequest"`
	CountryCallingCode         int    `json:"countryCallingCode"`
	Phone                      string `json:"phone"`
	ClientValidationFail       bool   `json:"clientValidationFail"`
	Locale                     string `json:"locale"`
	DateTimeTermsAndConditions int64  `json:"dtTC"`
	DateTimePrivacyNotes       int64  `json:"dtPN"`
	DateTimeLegalAge           int64  `json:"dtLA"`
	DeviceModel                string `json:"deviceModel"`
	OsVersion                  string `json:"osVersion"`
	Android                    bool   `json:"android"`
}

func (req StartReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type VerifyReq struct {
	WarmUpRequest    bool   `json:"warmUpRequest"`
	SessionId        string `json:"sessionId"`
	VerificationCode string `json:"verificationCode"`
}

func (req VerifyReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type VerifyResp struct {
	BaseResponse
	AccessToken         string `json:"accessToken"`
	AccountAlreadyExist bool   `json:"accountAlreadyExist"`
}

func (resp VerifyResp) GoString() string {
	return fmt.Sprintf("%#v", resp)
}

type CreateReq struct {
	WarmUpRequest bool   `json:"warmUpRequest"`
	AccessToken   string `json:"accessToken"`
	YearOfBirth   int    `json:"yearOfBirth"`
	Sex           string `json:"sex"`
}

func (req CreateReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type InternalGetUserIdReq struct {
	WarmUpRequest bool   `json:"warmUpRequest"`
	AccessToken   string `json:"accessToken"`
	BuildNum      int    `json:"buildNum"`
	IsItAndroid   bool   `json:"isItAndroid"`
}

func (req InternalGetUserIdReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type InternalGetUserIdResp struct {
	BaseResponse
	UserId         string `json:"userId"`
	IsUserReported bool   `json:"isUserReported"`
}

func (resp InternalGetUserIdResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

type UpdateSettingsReq struct {
	WarmUpRequest       bool   `json:"warmUpRequest"`
	AccessToken         string `json:"accessToken"`
	SafeDistanceInMeter int    `json:"safeDistanceInMeter"` // 0 (default for men) || 10 (default for women)
	PushMessages        bool   `json:"pushMessages"`        // true (default for men) || false (default for women)
	PushMatches         bool   `json:"pushMatches"`         // true (default)
	PushLikes           string `json:"pushLikes"`           //EVERY (default for men) || 10_NEW (default for women) || 100_NEW || NONE
}

func (req UpdateSettingsReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type GetSettingsResp struct {
	BaseResponse
	SafeDistanceInMeter int    `json:"safeDistanceInMeter"` // 0 (default for men) || 10 (default for women)
	PushMessages        bool   `json:"pushMessages"`        // true (default for men) || false (default for women)
	PushMatches         bool   `json:"pushMatches"`         // true (default)
	PushLikes           string `json:"pushLikes"`           //EVERY (default for men) || 10_NEW (default for women) || 100_NEW || NONE
}

func (resp GetSettingsResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

type LogoutReq struct {
	WarmUpRequest bool   `json:"warmUpRequest"`
	AccessToken   string `json:"accessToken"`
}

func (req LogoutReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type GetPresignUrlReq struct {
	WarmUpRequest bool   `json:"warmUpRequest"`
	AccessToken   string `json:"accessToken"`
	Extension     string `json:"extension"`
	ClientPhotoId string `json:"clientPhotoId"`
}

func (req GetPresignUrlReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type GetPresignUrlResp struct {
	BaseResponse
	Uri           string `json:"uri"`
	OriginPhotoId string `json:"originPhotoId"`
	ClientPhotoId string `json:"clientPhotoId"`
}

func (resp GetPresignUrlResp) GoString() string {
	return fmt.Sprintf("%#v", resp)
}

type MakePresignUrlInternalReq struct {
	WarmUpRequest bool   `json:"warmUpRequest"`
	Bucket        string `json:"bucket"`
	Key           string `json:"key"`
}

func (req MakePresignUrlInternalReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type MakePresignUrlInternalResp struct {
	Uri string `json:"uri"`
}

func (resp MakePresignUrlInternalResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

type GetOwnPhotosResp struct {
	BaseResponse
	Photos []OwnPhoto `json:"photos"`
}

func (resp GetOwnPhotosResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

type OwnPhoto struct {
	PhotoId       string `json:"photoId"`
	PhotoUri      string `json:"photoUri"`
	Likes         int    `json:"likes"`
	OriginPhotoId string `json:"originPhotoId"`
}

func (obj OwnPhoto) String() string {
	return fmt.Sprintf("%#v", obj)
}

type DeletePhotoReq struct {
	WarmUpRequest bool   `json:"warmUpRequest"`
	AccessToken   string `json:"accessToken"`
	PhotoId       string `json:"photoId"`
}

func (req DeletePhotoReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type Profile struct {
	UserId string  `json:"userId"`
	Photos []Photo `json:"photos"`
}

func (p Profile) String() string {
	return fmt.Sprintf("%#v", p)
}

type Photo struct {
	PhotoId  string `json:"photoId"`
	PhotoUri string `json:"photoUri"`
}

func (p Photo) String() string {
	return fmt.Sprintf("%#v", p)
}

type GetNewFacesResp struct {
	BaseResponse
	WarmUpRequest bool      `json:"warmUpRequest"`
	Profiles      []Profile `json:"profiles"`
}

func (resp GetNewFacesResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

type FacesWithUrlResp struct {
	//contains userId_photoId like a key and photoUrl like a value
	UserIdPhotoIdKeyUrlMap map[string]string `json:"urlPhotoMap"`
}

func (resp FacesWithUrlResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

type ActionReq struct {
	AccessToken string   `json:"accessToken"`
	Actions     []Action `json:"actions"`
}

func (req ActionReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type Action struct {
	SourceFeed    string `json:"sourceFeed"`
	ActionType    string `json:"actionType"`
	TargetPhotoId string `json:"targetPhotoId"`
	TargetUserId  string `json:"targetUserId"`
	LikeCount     int    `json:"likeCount"`
	ViewCount     int    `json:"viewCount"`
	ViewTimeSec   int    `json:"viewTimeSec"`
	ActionTime    int    `json:"actionTime"`
}

func (req Action) String() string {
	return fmt.Sprintf("%#v", req)
}
