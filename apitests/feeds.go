package apitests

import (
	"github.com/aws/aws-lambda-go/lambdacontext"
	"../apimodel"
	"github.com/satori/go.uuid"
)

func GetNewFacesWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	token := CreateUser("", "female", lc)
	resp := apimodel.GetNewFaces(token, apimodel.PhotoResolution480x640, false, lc)
	if resp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /get_new_faces with old build num")
	}
}

func GetNewFacesWithOldAccessToken(lc *lambdacontext.LambdaContext) {
	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	phone := uuid.String()
	token := CreateUser(phone, "male", lc)
	CreateUser(phone, "male", lc)
	resp := apimodel.GetNewFaces(token, apimodel.PhotoResolution480x640, true, lc)
	if resp.ErrorCode != "InvalidAccessTokenClientError" {
		panic("there is no InvalidAccessTokenClientError when call /get_new_faces with old build num")
	}
}

func GetNewFacesWithWrongResolution(lc *lambdacontext.LambdaContext) {
	token := CreateUser("", "female", lc)
	resp := apimodel.GetNewFaces(token, apimodel.PhotoResolution480x640+"sdf", true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /get_new_faces with wrong resolution")
	}
}
