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

func UpdateSettingsWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	resp := apimodel.UpdateUserSettings(token, 10, true, true, "EVERY", false, lc)
	if resp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /update_settings with old build num")
	}
}

func UpdateSettingsWithOldToken(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	apimodel.Logout(token, true, lc)
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

	apimodel.Anlogger.Debugf(lc, "test_auth.go : successfully create user")

	return baseResp.AccessToken
}
