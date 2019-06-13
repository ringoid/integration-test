package apitests

import (
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/ringoid/integration-test/apimodel"
	"fmt"
	"math/rand"
)

func CreateUserProfileWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	resp := apimodel.CreateUserProfile(1982, "male", false, lc)
	if resp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /create_profile with old build num")
	}
}

func CreateUserProfileWithWrongYearOfBirth(lc *lambdacontext.LambdaContext) {
	resp := apimodel.CreateUserProfile(0, "male", true, lc)
	if resp.ErrorCode != "WrongYearOfBirthClientError" {
		panic("there is no WrongYearOfBirthClientError when call /create_profile with wrong year of birth")
	}
}

func CreateUserProfileWithWrongSex(lc *lambdacontext.LambdaContext) {
	resp := apimodel.CreateUserProfile(1982, "malefemale", true, lc)
	if resp.ErrorCode != "WrongSexClientError" {
		panic("there is no WrongSexClientError when call /create_profile with wrong sex")
	}
}

func DeleteUserProfileWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	resp := apimodel.DeleteUserProfile(token, false, lc)
	if resp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /delete with old build num")
	}
}

func DeleteUserProfileWithWrongAccessToken(lc *lambdacontext.LambdaContext) {
	token := CreateUser("male", lc)
	DeleteUser(token, lc)
	resp := apimodel.DeleteUserProfile(token, true, lc)
	if resp.ErrorCode != "InvalidAccessTokenClientError" {
		panic("there is no InvalidAccessTokenClientError when call /delete with old token")
	}
}

func GetSettingsWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	token := CreateUser("male", lc)
	resp := apimodel.GetUserSettings(token, false, lc)
	if resp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /get_settings with old build num")
	}
}

func GetSettingsWithOldToken(lc *lambdacontext.LambdaContext) {
	token := CreateUser("male", lc)
	DeleteUser(token, lc)
	resp := apimodel.GetUserSettings(token, true, lc)
	if resp.ErrorCode != "InvalidAccessTokenClientError" {
		panic("there is no InvalidAccessTokenClientError when call /get_settings with old token")
	}
}

func UpdateSettingsWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	resp := apimodel.UpdateUserSettings(token, 10, true, true, "EVERY", false, lc)
	if resp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /update_settings with old build num")
	}
}

func UpdateSettingsWithOldToken(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	DeleteUser(token, lc)
	resp := apimodel.UpdateUserSettings(token, 10, true, true, "EVERY", true, lc)
	if resp.ErrorCode != "InvalidAccessTokenClientError" {
		panic("there is no InvalidAccessTokenClientError when call /update_settings with old token")
	}
}

func UpdateSettingsWithWrongDistance(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	resp := apimodel.UpdateUserSettings(token, -1, true, true, "EVERY", true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /update_settings with wrong distance")
	}
}

func UpdateSettingsWithWrongPushLikes(lc *lambdacontext.LambdaContext) {
	token := CreateUser("male", lc)
	resp := apimodel.UpdateUserSettings(token, 10, true, true, "EVERY11", true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /update_settings with wrong push likes")
	}
}

func CreateUser(sex string, lc *lambdacontext.LambdaContext) string {
	baseResp := apimodel.CreateUserProfile(1980+rand.Intn(20), sex, true, lc)
	if len(baseResp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error create user profile, error code %s", baseResp.ErrorCode))
	}

	apimodel.Anlogger.Debugf(lc, "auth.go : successfully create user")

	return baseResp.AccessToken
}

func UpdateProfile(accessToken string, isItCat bool,
	property, transport, income, height, educationLevel, hairColor, children int,
	lc *lambdacontext.LambdaContext) {

	resp := apimodel.UpdateUserProfile(accessToken, isItCat, property, transport, income, height, educationLevel, hairColor, children, true, lc)
	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error update user profile, error code %s", resp.ErrorCode))
	}

	apimodel.Anlogger.Debugf(lc, "auth.go : successfully update user profile")
}

func DeleteUser(token string, lc *lambdacontext.LambdaContext) {
	baseResp := apimodel.DeleteUserProfile(token, true, lc)
	if len(baseResp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error delete user profile, error code %s", baseResp.ErrorCode))
	}

	apimodel.Anlogger.Debugf(lc, "auth.go : successfully delete user")
}

func AuthTest(lc *lambdacontext.LambdaContext) {
	token := CreateUser("male", lc)
	getSettingResp := apimodel.GetUserSettings(token, true, lc)
	if len(getSettingResp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error when call /get_settings, token [%s], error code %s",
			token, getSettingResp.ErrorCode))
	}

	if getSettingResp.SafeDistanceInMeter != 0 ||
		getSettingResp.PushLikes != "EVERY" ||
		!getSettingResp.PushMessages ||
		!getSettingResp.PushMatches {
		panic(fmt.Sprintf("error with default setting for men, resp %v", getSettingResp))
	}

	updateSetResp := apimodel.UpdateUserSettings(token, 100, false, false, "10_NEW", true, lc)
	if len(updateSetResp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error when call /update_settings, token [%s], error code %s",
			token, updateSetResp.ErrorCode))
	}

	getSettingResp = apimodel.GetUserSettings(token, true, lc)
	if len(getSettingResp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error when call /get_settings, token [%s], error code %s",
			token, getSettingResp.ErrorCode))
	}

	if getSettingResp.SafeDistanceInMeter != 100 ||
		getSettingResp.PushLikes != "10_NEW" ||
		getSettingResp.PushMessages ||
		getSettingResp.PushMatches {
		panic(fmt.Sprintf("error after update settings, resp %v", getSettingResp))
	}

	//test for women
	token = CreateUser("female", lc)
	getSettingResp = apimodel.GetUserSettings(token, true, lc)
	if len(getSettingResp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error when call /get_settings, token [%s], error code %s",
			token, getSettingResp.ErrorCode))
	}

	if getSettingResp.SafeDistanceInMeter != 25 ||
		getSettingResp.PushLikes != "10_NEW" ||
		getSettingResp.PushMessages ||
		!getSettingResp.PushMatches {
		panic(fmt.Sprintf("error with default setting for women, resp %v", getSettingResp))
	}
}
