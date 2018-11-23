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
	apimodel.Logout(token, true, lc)
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
