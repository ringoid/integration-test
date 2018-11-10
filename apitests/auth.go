package apitests

import (
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/ringoid/integration-test/apimodel"
	"time"
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
)

func StartVerificationWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	request := apimodel.StartReq{
		CountryCallingCode:         7,
		Phone:                      "9113106409",
		ClientValidationFail:       false,
		Locale:                     "ru",
		DateTimeLegalAge:           time.Now().Unix(),
		DateTimePrivacyNotes:       time.Now().Unix(),
		DateTimeTermsAndConditions: time.Now().Unix(),
		DeviceModel:                "test device",
		OsVersion:                  "test android",
		Android:                    true,
	}
	resp := apimodel.StartRealAuth(request, false, lc)
	if resp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /start_verification with old build num")
	}
}

func StartVerificationWithWrongCountryCode(lc *lambdacontext.LambdaContext) {
	request := apimodel.StartReq{
		CountryCallingCode:         0,
		Phone:                      "9113106409",
		ClientValidationFail:       false,
		Locale:                     "ru",
		DateTimeLegalAge:           time.Now().Unix(),
		DateTimePrivacyNotes:       time.Now().Unix(),
		DateTimeTermsAndConditions: time.Now().Unix(),
		DeviceModel:                "test device",
		OsVersion:                  "test android",
		Android:                    true,
	}
	resp := apimodel.StartRealAuth(request, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /start_verification with wrong country code")
	}
}

func StartVerificationWithWrongPhone(lc *lambdacontext.LambdaContext) {
	request := apimodel.StartReq{
		CountryCallingCode:         7,
		Phone:                      "",
		ClientValidationFail:       false,
		Locale:                     "ru",
		DateTimeLegalAge:           time.Now().Unix(),
		DateTimePrivacyNotes:       time.Now().Unix(),
		DateTimeTermsAndConditions: time.Now().Unix(),
		DeviceModel:                "test device",
		OsVersion:                  "test android",
		Android:                    true,
	}
	resp := apimodel.StartRealAuth(request, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /start_verification with empty phone")
	}
}

func StartVerificationWithWrongLocale(lc *lambdacontext.LambdaContext) {
	request := apimodel.StartReq{
		CountryCallingCode:         7,
		Phone:                      "9334506790",
		ClientValidationFail:       false,
		Locale:                     "",
		DateTimeLegalAge:           time.Now().Unix(),
		DateTimePrivacyNotes:       time.Now().Unix(),
		DateTimeTermsAndConditions: time.Now().Unix(),
		DeviceModel:                "test device",
		OsVersion:                  "test android",
		Android:                    true,
	}
	resp := apimodel.StartRealAuth(request, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /start_verification with wrong locale")
	}
}

func StartVerificationWithWrongDateTimeLegalAge(lc *lambdacontext.LambdaContext) {
	request := apimodel.StartReq{
		CountryCallingCode:         7,
		Phone:                      "9334506790",
		ClientValidationFail:       false,
		Locale:                     "ru",
		DateTimeLegalAge:           0,
		DateTimePrivacyNotes:       time.Now().Unix(),
		DateTimeTermsAndConditions: time.Now().Unix(),
		DeviceModel:                "test device",
		OsVersion:                  "test android",
		Android:                    true,
	}
	resp := apimodel.StartRealAuth(request, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /start_verification with wrong date time legal age")
	}
}

func StartVerificationWithWrongDateTimePrivacyNotes(lc *lambdacontext.LambdaContext) {
	request := apimodel.StartReq{
		CountryCallingCode:         7,
		Phone:                      "9334506790",
		ClientValidationFail:       false,
		Locale:                     "ru",
		DateTimeLegalAge:           time.Now().Unix(),
		DateTimePrivacyNotes:       0,
		DateTimeTermsAndConditions: time.Now().Unix(),
		DeviceModel:                "test device",
		OsVersion:                  "test android",
		Android:                    true,
	}
	resp := apimodel.StartRealAuth(request, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /start_verification with wrong date time privacy notes")
	}
}

func StartVerificationWithWrongDateTimeTermsAndConditions(lc *lambdacontext.LambdaContext) {
	request := apimodel.StartReq{
		CountryCallingCode:         7,
		Phone:                      "9334506790",
		ClientValidationFail:       false,
		Locale:                     "ru",
		DateTimeLegalAge:           time.Now().Unix(),
		DateTimePrivacyNotes:       time.Now().Unix(),
		DateTimeTermsAndConditions: 0,
		DeviceModel:                "test device",
		OsVersion:                  "test android",
		Android:                    true,
	}
	resp := apimodel.StartRealAuth(request, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /start_verification with wrong date time terms and conditional")
	}
}

func StartVerificationWithWrongDeviceModel(lc *lambdacontext.LambdaContext) {
	request := apimodel.StartReq{
		CountryCallingCode:         7,
		Phone:                      "9334506790",
		ClientValidationFail:       false,
		Locale:                     "ru",
		DateTimeLegalAge:           time.Now().Unix(),
		DateTimePrivacyNotes:       time.Now().Unix(),
		DateTimeTermsAndConditions: time.Now().Unix(),
		DeviceModel:                "",
		OsVersion:                  "test android",
		Android:                    true,
	}
	resp := apimodel.StartRealAuth(request, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /start_verification with wrong device model")
	}
}

func StartVerificationWithWrongOsVersion(lc *lambdacontext.LambdaContext) {
	request := apimodel.StartReq{
		CountryCallingCode:         7,
		Phone:                      "9334506790",
		ClientValidationFail:       false,
		Locale:                     "ru",
		DateTimeLegalAge:           time.Now().Unix(),
		DateTimePrivacyNotes:       time.Now().Unix(),
		DateTimeTermsAndConditions: time.Now().Unix(),
		DeviceModel:                "test model",
		OsVersion:                  "",
		Android:                    true,
	}
	resp := apimodel.StartRealAuth(request, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /start_verification with wrong os")
	}
}

func CompleteVerificationWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	request := apimodel.VerifyReq{
		SessionId:        "fake_session",
		VerificationCode: "fake_code",
	}

	resp := apimodel.CompleteRealAuth(request, false, lc)
	if resp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /complete_verification with old build num")
	}
}

func CompleteVerificationWithWrongSessionId(lc *lambdacontext.LambdaContext) {
	request := apimodel.VerifyReq{
		SessionId:        "",
		VerificationCode: "fake_code",
	}

	resp := apimodel.CompleteRealAuth(request, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /complete_verification with wrong session id")
	}
}

func CompleteVerificationWithWrongVerificationCode(lc *lambdacontext.LambdaContext) {
	request := apimodel.VerifyReq{
		SessionId:        "fake_session",
		VerificationCode: "",
	}

	resp := apimodel.CompleteRealAuth(request, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /complete_verification with wrong verification code")
	}
}

func CreateUserProfileWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	token := verifyUser("", lc)
	resp := apimodel.CreateUserProfile(token, 1982, "male", false, lc)
	if resp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /create_profile with old build num")
	}
}

func CreateUserProfileWithOldToken(lc *lambdacontext.LambdaContext) {
	token := verifyUser("1112220000", lc)
	verifyUser("1112220000", lc)

	resp := apimodel.CreateUserProfile(token, 1982, "male", true, lc)
	if resp.ErrorCode != "InvalidAccessTokenClientError" {
		panic("there is no InvalidAccessTokenClientError when call /create_profile with old token")
	}
}

func CreateUserProfileWithWrongYearOfBirth(lc *lambdacontext.LambdaContext) {
	token := verifyUser("", lc)
	resp := apimodel.CreateUserProfile(token, 0, "male", true, lc)
	if resp.ErrorCode != "WrongYearOfBirthClientError" {
		panic("there is no WrongYearOfBirthClientError when call /create_profile with wrong year of birth")
	}
}

func CreateUserProfileWithWrongSex(lc *lambdacontext.LambdaContext) {
	token := verifyUser("", lc)
	resp := apimodel.CreateUserProfile(token, 1982, "malefemale", true, lc)
	if resp.ErrorCode != "WrongSexClientError" {
		panic("there is no WrongSexClientError when call /create_profile with wrong sex")
	}
}

func UpdateSettingsBeforeCreateProfile(lc *lambdacontext.LambdaContext) {
	token := verifyUser("", lc)
	resp := apimodel.UpdateUserSettings(token, 10, true, true, "EVERY", true, lc)
	if resp.ErrorCode != "InvalidAccessTokenClientError" {
		panic("there is no InvalidAccessTokenClientError when call /update_settings before create profile")
	}
}

func UpdateSettingsWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	token := verifyUser("", lc)
	apimodel.CreateUserProfile(token, 1982, "male", false, lc)
	resp := apimodel.UpdateUserSettings(token, 10, true, true, "EVERY", false, lc)
	if resp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /update_settings with old build num")
	}
}

func UpdateSettingsWithOldToken(lc *lambdacontext.LambdaContext) {
	token := verifyUser("0101010101", lc)
	apimodel.CreateUserProfile(token, 1982, "male", false, lc)
	verifyUser("0101010101", lc)
	resp := apimodel.UpdateUserSettings(token, 10, true, true, "EVERY", true, lc)
	if resp.ErrorCode != "InvalidAccessTokenClientError" {
		panic("there is no InvalidAccessTokenClientError when call /update_settings with old token")
	}
}

func UpdateSettingsWithWrongDistance(lc *lambdacontext.LambdaContext) {
	token := verifyUser("", lc)
	apimodel.CreateUserProfile(token, 1982, "male", false, lc)
	resp := apimodel.UpdateUserSettings(token, -1, true, true, "EVERY", true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /update_settings with wrong distance")
	}
}

func UpdateSettingsWithWrongPushLikes(lc *lambdacontext.LambdaContext) {
	token := verifyUser("", lc)
	apimodel.CreateUserProfile(token, 1982, "male", false, lc)
	resp := apimodel.UpdateUserSettings(token, 10, true, true, "EVERY11", true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /update_settings with wrong push likes")
	}
}

func verifyUser(phone string, lc *lambdacontext.LambdaContext) string {
	apimodel.Anlogger.Debugf(lc, "test_auth.go : start create user")
	if len(phone) == 0 {
		uuid, err := uuid.NewV4()
		if err != nil {
			panic(err)
		}
		phone = uuid.String()
	}
	authResp := apimodel.Start(7, phone, lc)
	if len(authResp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error start registration, error code %s", authResp.ErrorCode))
	}

	verifyResp := apimodel.Complete(authResp.SessionId, "fake_code", lc)
	if len(verifyResp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error complete registration, error code %s", verifyResp.ErrorCode))
	}

	return verifyResp.AccessToken
}

func CreateUser(phone, sex string, lc *lambdacontext.LambdaContext) string {
	token := verifyUser(phone, lc)

	baseResp := apimodel.CreateUserProfile(token, 1980+rand.Intn(20), sex, true, lc)
	if len(baseResp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error create user profile, error code %s", baseResp.ErrorCode))
	}

	apimodel.Anlogger.Debugf(lc, "test_auth.go : successfully create user")

	return token
}
