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
type CreateReq struct {
	WarmUpRequest              bool   `json:"warmUpRequest"`
	YearOfBirth                int    `json:"yearOfBirth"`
	Sex                        string `json:"sex"`
	Locale                     string `json:"locale"`
	DateTimeTermsAndConditions int64  `json:"dtTC"`
	DateTimePrivacyNotes       int64  `json:"dtPN"`
	DateTimeLegalAge           int64  `json:"dtLA"`
	DeviceModel                string `json:"deviceModel"`
	OsVersion                  string `json:"osVersion"`
}

func (req CreateReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type CreateResp struct {
	BaseResponse
	AccessToken string `json:"accessToken"`
}

func (resp CreateResp) String() string {
	return fmt.Sprintf("%#v", resp)
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

type UpdateProfileInfoReq struct {
	AccessToken    string `json:"accessToken"`
	Property       int    `json:"property"`
	Transport      int    `json:"transport"`
	Income         int    `json:"income"`
	Height         int    `json:"height"`
	EducationLevel int    `json:"educationLevel"`
	HairColor      int    `json:"hairColor"`
	Children       int    `json:"children"`
	Name           string `json:"name"`
	JobTitle       string `json:"jobTitle"`
	Company        string `json:"company"`
	EducationText  string `json:"education"`
	About          string `json:"about"`
	Instagram      string `json:"instagram"`
	TikTok         string `json:"tikTok"`
	WhereLive      string `json:"whereLive"`
	WhereFrom      string `json:"whereFrom"`
	StatusText     string `json:"statusText"`
}

func (req UpdateProfileInfoReq) String() string {
	return fmt.Sprintf("%#v", req)
}

type GetSettingsResp struct {
	BaseResponse
	SafeDistanceInMeter int    `json:"safeDistanceInMeter"` // 0 (default for men) || 25 (default for women)
	PushMessages        bool   `json:"pushMessages"`        // true (default for men) || false (default for women)
	PushMatches         bool   `json:"pushMatches"`         // true (default)
	PushLikes           string `json:"pushLikes"`           //EVERY (default for men) || 10_NEW (default for women) || 100_NEW || NONE
}

func (resp GetSettingsResp) String() string {
	return fmt.Sprintf("%#v", resp)
}

type DeleteReq struct {
	WarmUpRequest bool   `json:"warmUpRequest"`
	AccessToken   string `json:"accessToken"`
}

func (req DeleteReq) String() string {
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
	UserId                      string    `json:"userId"`
	DefaultSortingOrderPosition int       `json:"defaultSortingOrderPosition"`
	Unseen                      bool      `json:"notSeen"`
	Photos                      []Photo   `json:"photos"`
	Messages                    []Message `json:"messages"`
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

type Message struct {
	WasYouSender bool   `json:"wasYouSender"`
	Text         string `json:"text"`
}

type GetNewFacesResp struct {
	BaseResponse
	WarmUpRequest      bool      `json:"warmUpRequest"`
	Profiles           []Profile `json:"profiles"`
	RepeatRequestAfter int       `json:"repeatRequestAfter"`
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
	SourceFeed      string  `json:"sourceFeed"`
	ActionType      string  `json:"actionType"`
	TargetPhotoId   string  `json:"targetPhotoId"`
	TargetUserId    string  `json:"targetUserId"`
	LikeCount       int     `json:"likeCount"`
	ViewCount       int     `json:"viewCount"`
	ViewTimeMillis  int64   `json:"viewTimeMillis"`
	Text            string  `json:"text"`
	Lat             float64 `json:"lat"`
	Lon             float64 `json:"lon"`
	ActionTime      int64   `json:"actionTime"`
	ClientMessageId string  `json:"clientMsgId"`
}

func (req Action) String() string {
	return fmt.Sprintf("%#v", req)
}

type LMMFeedResp struct {
	BaseResponse
	LikesYou           []Profile `json:"likesYou"`
	Matches            []Profile `json:"matches"`
	Messages           []Profile `json:"messages"`
	RepeatRequestAfter int       `json:"repeatRequestAfter"`
}

func (resp LMMFeedResp) String() string {
	return fmt.Sprintf("%#v", resp)
}
