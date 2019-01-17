package apimodel

import (
	"github.com/aws/aws-sdk-go/service/lambda"
	"os"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"encoding/json"
	"time"
	"net/http"
	"bytes"
	"io/ioutil"
	"strings"
	"github.com/ringoid/commons"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var Anlogger *commons.Logger

var clientLambda *lambda.Lambda
var authApiEndpoint string
var imageApiEndpoint string
var actionsApiEndpoint string
var feedsApiEndpoint string
var cleanAuthDbFunctionName string
var cleanImageDbFunctionName string
var cleanRelationshipsDbFunctionName string
var secretWord string

var AwsDynamoDB *dynamodb.DynamoDB
var BotsTableName string

func init() {
	var env string
	var ok bool
	var papertrailAddress string
	var err error
	var awsSession *session.Session

	env, ok = os.LookupEnv("ENV")
	if !ok {
		fmt.Printf("lambda-initialization : common_action.go : env can not be empty ENV\n")
		os.Exit(1)
	}
	fmt.Printf("lambda-initialization : common_action.go : start with ENV = [%s]\n", env)

	papertrailAddress, ok = os.LookupEnv("PAPERTRAIL_LOG_ADDRESS")
	if !ok {
		fmt.Printf("lambda-initialization : common_action.go : env can not be empty PAPERTRAIL_LOG_ADDRESS\n")
		os.Exit(1)
	}
	fmt.Printf("lambda-initialization : common_action.go : start with PAPERTRAIL_LOG_ADDRESS = [%s]\n", papertrailAddress)

	Anlogger, err = commons.New(papertrailAddress, fmt.Sprintf("%s-%s", env, "integration-test"))
	if err != nil {
		fmt.Errorf("lambda-initialization : common_action.go : error during startup : %v\n", err)
		os.Exit(1)
	}
	Anlogger.Debugf(nil, "lambda-initialization : common_action.go : logger was successfully initialized")

	authApiEndpoint, ok = os.LookupEnv("AUTH_API_ENDPOINT")
	if !ok {
		Anlogger.Fatalf(nil, "lambda-initialization : common_action.go : env can not be empty AUTH_API_ENDPOINT")
	}
	Anlogger.Debugf(nil, "lambda-initialization : common_action.go : start with AUTH_API_ENDPOINT = [%s]", authApiEndpoint)

	imageApiEndpoint, ok = os.LookupEnv("IMAGE_API_ENDPOINT")
	if !ok {
		Anlogger.Fatalf(nil, "lambda-initialization : common_action.go : env can not be empty IMAGE_API_ENDPOINT")
	}
	Anlogger.Debugf(nil, "lambda-initialization : common_action.go : start with IMAGE_API_ENDPOINT = [%s]", imageApiEndpoint)

	actionsApiEndpoint, ok = os.LookupEnv("ACTIONS_API_ENDPOINT")
	if !ok {
		Anlogger.Fatalf(nil, "lambda-initialization : common_action.go : env can not be empty ACTIONS_API_ENDPOINT")
	}
	Anlogger.Debugf(nil, "lambda-initialization : common_action.go : start with ACTIONS_API_ENDPOINT = [%s]", actionsApiEndpoint)

	feedsApiEndpoint, ok = os.LookupEnv("FEEDS_API_ENDPOINT")
	if !ok {
		Anlogger.Fatalf(nil, "lambda-initialization : common_action.go : env can not be empty FEEDS_API_ENDPOINT")
	}
	Anlogger.Debugf(nil, "lambda-initialization : common_action.go : start with FEEDS_API_ENDPOINT = [%s]", feedsApiEndpoint)

	cleanAuthDbFunctionName, ok = os.LookupEnv("CLEAN_AUTH_DB_FUNCTION_NAME")
	if !ok {
		Anlogger.Fatalf(nil, "lambda-initialization : common_action.go : env can not be empty CLEAN_AUTH_DB_FUNCTION_NAME")
	}
	Anlogger.Debugf(nil, "lambda-initialization : common_action.go : start with CLEAN_AUTH_DB_FUNCTION_NAME = [%s]", cleanAuthDbFunctionName)

	cleanImageDbFunctionName, ok = os.LookupEnv("CLEAN_IMAGE_DB_FUNCTION_NAME")
	if !ok {
		Anlogger.Fatalf(nil, "lambda-initialization : common_action.go : env can not be empty CLEAN_IMAGE_DB_FUNCTION_NAME")
	}
	Anlogger.Debugf(nil, "lambda-initialization : common_action.go : start with CLEAN_IMAGE_DB_FUNCTION_NAME = [%s]", cleanImageDbFunctionName)

	BotsTableName, ok = os.LookupEnv("BOTS_TABLE_NAME")
	if !ok {
		Anlogger.Fatalf(nil, "lambda-initialization : common_action.go : env can not be empty BOTS_TABLE_NAME")
	}
	Anlogger.Debugf(nil, "lambda-initialization : common_action.go : start with BOTS_TABLE_NAME = [%s]", BotsTableName)

	cleanRelationshipsDbFunctionName, ok = os.LookupEnv("CLEAN_RELATIONSHIPS_DB_FUNCTION_NAME")
	if !ok {
		Anlogger.Fatalf(nil, "lambda-initialization : common_action.go : env can not be empty CLEAN_RELATIONSHIPS_DB_FUNCTION_NAME")
	}
	Anlogger.Debugf(nil, "lambda-initialization : common_action.go : start with CLEAN_RELATIONSHIPS_DB_FUNCTION_NAME = [%s]", cleanRelationshipsDbFunctionName)

	awsSession, err = session.NewSession(aws.NewConfig().
		WithRegion(Region).WithMaxRetries(MaxRetries).
		WithLogger(aws.LoggerFunc(func(args ...interface{}) { Anlogger.AwsLog(args) })).WithLogLevel(aws.LogOff))
	if err != nil {
		Anlogger.Fatalf(nil, "lambda-initialization : common_action.go : error during initialization : %v", err)
	}
	Anlogger.Debugf(nil, "lambda-initialization : common_action.go : aws session was successfully initialized")

	clientLambda = lambda.New(awsSession)
	Anlogger.Debugf(nil, "lambda-initialization : common_action.go : lambda client was successfully initialized")

	secretWord = commons.GetSecret(fmt.Sprintf(commons.SecretWordKeyBase, env), commons.SecretWordKeyName, awsSession, Anlogger, nil)

	AwsDynamoDB = dynamodb.New(awsSession)
	Anlogger.Debugf(nil, "lambda-initialization : common_action.go : dynamodb client was successfully initialized")
}

func CleanAllDB(lc *lambdacontext.LambdaContext) {
	Anlogger.Warnf(lc, "common_action.go : start clean all DB")
	_, err := clientLambda.Invoke(&lambda.InvokeInput{FunctionName: aws.String(cleanAuthDbFunctionName),})
	if err != nil {
		panic(err)
	}
	Anlogger.Warnf(lc, "common_action.go : successfully clean %s", cleanAuthDbFunctionName)
	_, err = clientLambda.Invoke(&lambda.InvokeInput{FunctionName: aws.String(cleanImageDbFunctionName),})
	if err != nil {
		panic(err)
	}
	Anlogger.Warnf(lc, "common_action.go : successfully clean %s", cleanImageDbFunctionName)
	_, err = clientLambda.Invoke(&lambda.InvokeInput{FunctionName: aws.String(cleanRelationshipsDbFunctionName),})
	if err != nil {
		panic(err)
	}
	Anlogger.Warnf(lc, "common_action.go : successfully clean %s", cleanRelationshipsDbFunctionName)
	Anlogger.Warnf(lc, "common_action.go : successfully end clean all DB")
}

//Auth service
func CreateUserProfile(yearOfBirth int, sex string, useValidBuildNum bool, lc *lambdacontext.LambdaContext) CreateResp {
	Anlogger.Debugf(lc, "common_action.go : create user profile, yearOfBirth [%d], sex [%s], useValidBuildNum [%v]",
		yearOfBirth, sex, useValidBuildNum)

	request := CreateReq{
		YearOfBirth:                yearOfBirth,
		Sex:                        sex,
		Locale:                     "ru",
		DateTimeLegalAge:           time.Now().Unix(),
		DateTimePrivacyNotes:       time.Now().Unix(),
		DateTimeTermsAndConditions: time.Now().Unix(),
		DeviceModel:                "test device",
		OsVersion:                  "test android",
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	respBody := makePostRequest(authApiEndpoint, jsonBody, "/create_profile", useValidBuildNum)

	response := CreateResp{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}

	Anlogger.Infof(lc, "common_action.go : successfully call create user profile, return response %v", response)
	return response
}

func DeleteUserProfile(accessToken string, useValidBuildNum bool, lc *lambdacontext.LambdaContext) BaseResponse {
	Anlogger.Debugf(lc, "common_action.go : delete user profile, use valid build num [%d], accessToken [%s]",
		useValidBuildNum, accessToken)

	request := DeleteReq{
		AccessToken: accessToken,
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	respBody := makePostRequest(authApiEndpoint, jsonBody, "/delete", useValidBuildNum)

	response := BaseResponse{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}

	Anlogger.Infof(lc, "common_action.go : successfully call delete user profile, return response %v", response)
	return response
}

func UpdateUserSettings(accessToken string, safeDistanceInMeter int, pushMessages, pushMatches bool,
	pushLikes string, useValidBuildNum bool, lc *lambdacontext.LambdaContext) BaseResponse {
	Anlogger.Debugf(lc, "common_action.go : update user's settings, token [%s], safeDistanceInMeter [%d], "+
		"pushMessages [%d], pushMatches [%d], pushLikes [%s], useValidBuildNum [%v]",
		accessToken, safeDistanceInMeter, pushMessages, pushMatches, pushLikes, useValidBuildNum)

	request := UpdateSettingsReq{
		AccessToken:         accessToken,
		SafeDistanceInMeter: safeDistanceInMeter,
		PushMessages:        pushMessages,
		PushMatches:         pushMatches,
		PushLikes:           pushLikes,
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	respBody := makePostRequest(authApiEndpoint, jsonBody, "/update_settings", useValidBuildNum)

	response := BaseResponse{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}

	Anlogger.Infof(lc, "common_action.go : successfully call update user's settings, return response %v", response)
	return response
}

func GetUserSettings(accessToken string, useValidBuildNum bool, lc *lambdacontext.LambdaContext) GetSettingsResp {
	Anlogger.Debugf(lc, "common_action.go : get user's settings, token [%s], useValidBuildNum [%v]", accessToken, useValidBuildNum)

	params := make(map[string]string)
	params["accessToken"] = accessToken
	respBody := makeGetRequest(authApiEndpoint, params, "/get_settings", useValidBuildNum)
	response := GetSettingsResp{}
	err := json.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}

	Anlogger.Infof(lc, "common_action.go : successfully call get user's settings, return response %v", response)
	return response
}

func makeGetRequest(baseApi string, params map[string]string, urlPart string, useValidBuildNum bool) []byte {
	url := baseApi + urlPart + "?"
	for k, v := range params {
		url += fmt.Sprintf("%s=%s&", k, v)
	}
	index := strings.LastIndex(url, "&")
	url = url[:index]

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("x-ringoid-android-buildnum", "1000")
	if !useValidBuildNum {
		req.Header.Set("x-ringoid-android-buildnum", "0")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	httpResponse, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		panic(fmt.Sprintf("error call [%s], status code [%d]", baseApi+urlPart, httpResponse.StatusCode))
	}

	respBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		panic(err)
	}

	return respBody
}

func makePostRequest(baseApi string, jsonBody []byte, urlPart string, useValidBuildNum bool) []byte {
	req, err := http.NewRequest("POST", baseApi+urlPart, bytes.NewReader(jsonBody))
	if err != nil {
		panic(err)
	}
	req.Header.Set("x-ringoid-android-buildnum", "1000")
	if !useValidBuildNum {
		req.Header.Set("x-ringoid-android-buildnum", "0")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	httpResponse, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		panic(fmt.Sprintf("error call [%s], status code [%d]", baseApi+urlPart, httpResponse.StatusCode))
	}

	respBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		panic(err)
	}

	return respBody
}

//Image service

func GenerateImage(isItMan bool, text string) []byte {
	white := "ffffff"
	black := "000000"
	blue := "0011ff"
	//background, text color, text
	urlStr := "https://dummyimage.com/900/%s/%s.jpg&text=%s"
	if isItMan {
		urlStr = fmt.Sprintf(urlStr, blue, white, text)
	} else {
		urlStr = fmt.Sprintf(urlStr, black, white, text)
	}

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	httpResponse, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		panic(fmt.Sprintf("error call generate image, status code [%d]", httpResponse.StatusCode))
	}

	respBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		panic(err)
	}

	return respBody
}

func MakePutRequestWithContent(url string, source []byte) {
	request, err := http.NewRequest("PUT", url, bytes.NewReader(source))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	httpResponse, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer httpResponse.Body.Close()

	_, err = ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		panic(err)
	}

	if httpResponse.StatusCode != 200 {
		panic(fmt.Sprintf("error making put request with content, url [%s],  status code [%d]", url, httpResponse.StatusCode))
	}
}

func GetPresignUrl(accessToken string, useValidBuildNum bool, lc *lambdacontext.LambdaContext) GetPresignUrlResp {
	Anlogger.Debugf(lc, "common_action.go : get presigned url, token [%s], useValidBuildNum [%v]", accessToken, useValidBuildNum)

	request := GetPresignUrlReq{
		AccessToken:   accessToken,
		Extension:     "jpg",
		ClientPhotoId: "fakeClientId",
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	respBody := makePostRequest(imageApiEndpoint, jsonBody, "/get_presigned", useValidBuildNum)

	response := GetPresignUrlResp{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}

	Anlogger.Infof(lc, "comon_action.go : successfully call get presigned url, return response %v", response)
	return response
}

func GetOwnPhotos(accessToken, resolution string, useValidBuildNum bool, lc *lambdacontext.LambdaContext) GetOwnPhotosResp {
	Anlogger.Debugf(lc, "common_action.go : get own photos, token [%s], resolution [%s], useValidBuildNum [%v]", accessToken, resolution, useValidBuildNum)

	params := make(map[string]string)
	params["accessToken"] = accessToken
	params["resolution"] = resolution
	respBody := makeGetRequest(imageApiEndpoint, params, "/get_own_photos", useValidBuildNum)
	response := GetOwnPhotosResp{}
	err := json.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}

	Anlogger.Infof(lc, "common_action.go : successfully call get own photos, return response %v", response)
	return response
}

func DeletePhoto(accessToken, photoId string, useValidBuildNum bool, lc *lambdacontext.LambdaContext) BaseResponse {
	Anlogger.Debugf(lc, "common_action.go : delete photo, token [%s], photoId [%s], useValidBuildNum [%v]", accessToken, photoId, useValidBuildNum)
	request := DeletePhotoReq{
		AccessToken: accessToken,
		PhotoId:     photoId,
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	respBody := makePostRequest(imageApiEndpoint, jsonBody, "/delete_photo", useValidBuildNum)

	response := BaseResponse{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}

	Anlogger.Infof(lc, "common_action.go : successfully call delete photo, return response %v", response)
	return response
}

func Actions(accessToken string, actions []Action, useValidBuildNum bool, lc *lambdacontext.LambdaContext) BaseResponse {
	Anlogger.Debugf(lc, "common_action.go : actions, token [%s], actions %v, valid build num [%v]", accessToken, actions, useValidBuildNum)
	request := ActionReq{
		AccessToken: accessToken,
		Actions:     actions,
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	respBody := makePostRequest(actionsApiEndpoint, jsonBody, "/actions", useValidBuildNum)

	response := BaseResponse{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}
	Anlogger.Infof(lc, "common_action.go : successfully call actions, return response %v", response)
	return response
}

func GetNewFaces(accessToken, resolution string, lastActionTime int64, useValidBuildNum bool, lc *lambdacontext.LambdaContext) GetNewFacesResp {
	Anlogger.Debugf(lc, "common_action.go : actions, token [%s], resolution [%s]", accessToken, resolution)
	params := make(map[string]string)
	params["accessToken"] = accessToken
	params["resolution"] = resolution
	params["lastActionTime"] = fmt.Sprintf("%v", lastActionTime)
	respBody := makeGetRequest(feedsApiEndpoint, params, "/get_new_faces", useValidBuildNum)
	response := GetNewFacesResp{}
	err := json.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}

	Anlogger.Infof(lc, "common_action.go : successfully call new faces, response %v", response)
	return response
}

func GetLMM(accessToken, resolution string, lastActionTime int64, useValidBuildNum bool, lc *lambdacontext.LambdaContext) LMMFeedResp {
	Anlogger.Debugf(lc, "common_action.go : llm, token [%s], resolution [%s], lastActionTime [%d], use valid build num [%v]",
		accessToken, resolution, lastActionTime, useValidBuildNum)
	params := make(map[string]string)
	params["accessToken"] = accessToken
	params["resolution"] = resolution
	params["lastActionTime"] = fmt.Sprintf("%v", lastActionTime)
	respBody := makeGetRequest(feedsApiEndpoint, params, "/get_lmm", useValidBuildNum)
	response := LMMFeedResp{}
	err := json.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}

	Anlogger.Infof(lc, "common_action.go : successfully call llm, response %v", response)
	return response
}

func UserId(token string, lc *lambdacontext.LambdaContext) string {
	userId, _, ok, errorStr := commons.DecodeToken(token, secretWord, Anlogger, lc)
	if !ok {
		panic(fmt.Sprintf("could not decode token : %s", errorStr))
	}
	return userId
}
