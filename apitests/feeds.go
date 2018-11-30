package apitests

import (
	"github.com/aws/aws-lambda-go/lambdacontext"
	"../apimodel"
)

func GetNewFacesWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	resp := apimodel.GetNewFaces(token, apimodel.PhotoResolution480x640, false, lc)
	if resp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /get_new_faces with old build num")
	}
}

func GetNewFacesWithOldAccessToken(lc *lambdacontext.LambdaContext) {
	token := CreateUser("male", lc)
	DeleteUser(token, lc)
	resp := apimodel.GetNewFaces(token, apimodel.PhotoResolution480x640, true, lc)
	if resp.ErrorCode != "InvalidAccessTokenClientError" {
		panic("there is no InvalidAccessTokenClientError when call /get_new_faces with old build num")
	}
}

func GetNewFacesWithWrongResolution(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	resp := apimodel.GetNewFaces(token, apimodel.PhotoResolution480x640+"sdf", true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /get_new_faces with wrong resolution")
	}
}

func GetLMMWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	resp := apimodel.GetLMM(token, apimodel.PhotoResolution480x640, 0, false, lc)
	if resp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /get_lmm with old build num")
	}
}

func GetLMMWithOldAccessToken(lc *lambdacontext.LambdaContext) {
	token := CreateUser("male", lc)
	DeleteUser(token, lc)
	resp := apimodel.GetLMM(token, apimodel.PhotoResolution480x640, 0, true, lc)
	if resp.ErrorCode != "InvalidAccessTokenClientError" {
		panic("there is no InvalidAccessTokenClientError when call /get_lmm with old build num")
	}
}

func GetLMMWithWrongResolution(lc *lambdacontext.LambdaContext) {
	token := CreateUser("female", lc)
	resp := apimodel.GetLMM(token, apimodel.PhotoResolution480x640+"sdf", 0, true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /get_lmm with wrong resolution")
	}
}
