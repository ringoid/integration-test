package apitests

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"../apimodel"
	"github.com/satori/go.uuid"
	"time"
)

func GetPresignWithOldBuildNumTest(lc *lambdacontext.LambdaContext) {
	token := CreateUser("male", lc)
	getPresignResp := apimodel.GetPresignUrl(token, false, lc)
	if getPresignResp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /get_presigned with old build num")
	}
}

func GetPresignWithOldAccessTokenTest(lc *lambdacontext.LambdaContext) {
	token := CreateUser("male", lc)
	DeleteUser(token, lc)
	getPresignResp := apimodel.GetPresignUrl(token, true, lc)
	if getPresignResp.ErrorCode != "InvalidAccessTokenClientError" {
		panic("there is no InvalidAccessTokenClientError when call /get_presigned with old token")
	}
}

func GetOwnPhotosWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	token := CreateUser("male", lc)
	getOwnPhotoResp := apimodel.GetOwnPhotos(token, apimodel.PhotoResolution480x640, false, lc)
	if getOwnPhotoResp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /get_own_photos with old build num")
	}
}

func GetOwnPhotosWithOldToken(lc *lambdacontext.LambdaContext) {
	token := CreateUser("male", lc)
	DeleteUser(token, lc)
	getOwnPhotoResp := apimodel.GetOwnPhotos(token, apimodel.PhotoResolution480x640, true, lc)
	if getOwnPhotoResp.ErrorCode != "InvalidAccessTokenClientError" {
		panic("there is no InvalidAccessTokenClientError when call /get_own_photos with old token")
	}
}

func GetOwnPhotosWithWrongResolution(lc *lambdacontext.LambdaContext) {
	token := CreateUser("male", lc)
	getOwnPhotoResp := apimodel.GetOwnPhotos(token, apimodel.PhotoResolution480x640+"22", true, lc)
	if getOwnPhotoResp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /get_own_photos with wrong resolution")
	}
}

func DeletePhotoWithOldBuildNum(lc *lambdacontext.LambdaContext) {
	token := CreateUser("male", lc)
	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	resp := apimodel.DeletePhoto(token, uuid.String(), false, lc)
	if resp.ErrorCode != "TooOldAppVersionClientError" {
		panic("there is no TooOldAppVersionClientError when call /delete_photo with old build num")
	}
}

func DeletePhotoWithOldAccessToken(lc *lambdacontext.LambdaContext) {
	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	token := CreateUser("male", lc)
	DeleteUser(token, lc)
	resp := apimodel.DeletePhoto(token, uuid.String(), true, lc)
	if resp.ErrorCode != "InvalidAccessTokenClientError" {
		panic("there is no InvalidAccessTokenClientError when call /delete_photo with old build num")
	}
}

func DeletePhotoWithEmptyPhotoId(lc *lambdacontext.LambdaContext) {
	token := CreateUser("male", lc)
	resp := apimodel.DeletePhoto(token, "", true, lc)
	if resp.ErrorCode != "WrongParamsClientError" {
		panic("there is no WrongParamsClientError when call /delete_photo with empty photo id")
	}
}

func ImageTest(lc *lambdacontext.LambdaContext) {
	apimodel.Anlogger.Debugf(lc, "images.go : start image service complex test")
	token_1 := CreateUser("male", lc)
	//check that users doesn't have own photos
	resp := apimodel.GetOwnPhotos(token_1, apimodel.PhotoResolution480x640, true, lc)
	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error when call /get_own_photos, token [%s], resolution [%s], error code %s",
			token_1, apimodel.PhotoResolution480x640, resp.ErrorCode))
	}
	if len(resp.Photos) != 0 {
		panic(fmt.Sprintf("token_1 already have own photos, result %v", resp))
	}

	//now upload and check
	originPhotoId := UploadImage(token_1, "male", 1, 1, lc)

	resp = apimodel.GetOwnPhotos(token_1, apimodel.PhotoResolution480x640, true, lc)
	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error when call /get_own_photos, token [%s], resolution [%s], error code %s",
			token_1, apimodel.PhotoResolution480x640, resp.ErrorCode))
	}
	if len(resp.Photos) == 0 {
		apimodel.Anlogger.Warnf(lc, "there is no own photos immediately after uploading")
		time.Sleep(time.Second * 3)
	}

	resp = apimodel.GetOwnPhotos(token_1, apimodel.PhotoResolution480x640, true, lc)
	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error when call /get_own_photos, token [%s], resolution [%s], error code %s",
			token_1, apimodel.PhotoResolution480x640, resp.ErrorCode))
	}
	if len(resp.Photos) != 1 {
		panic(fmt.Sprintf("token_1 has wrong num of photos, result %v", resp))
	}

	photo := resp.Photos[0]
	if photo.OriginPhotoId != originPhotoId {
		panic(fmt.Sprintf("received originPhotoId doesn't match, originPhotoId [%s], resp %v",
			originPhotoId, resp))
	}

	if len(photo.PhotoId) == 0 || len(photo.PhotoUri) == 0 || photo.Likes != 0 {
		panic(fmt.Sprintf("received photo is in wrong state, resp %v", resp))
	}

	originPhotoId_2 := UploadImage(token_1, "male", 1, 2, lc)
	time.Sleep(time.Second * 3)

	resp = apimodel.GetOwnPhotos(token_1, apimodel.PhotoResolution480x640, true, lc)
	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error when call /get_own_photos, token [%s], resolution [%s], error code %s",
			token_1, apimodel.PhotoResolution480x640, resp.ErrorCode))
	}
	if len(resp.Photos) != 2 {
		panic(fmt.Sprintf("token_1 has wrong num of photos, result %v", resp))
	}

	//now check sorting order - last uploaded should be first
	photo_1 := resp.Photos[0]
	if photo_1.OriginPhotoId != originPhotoId_2 {
		panic(fmt.Sprintf("received originPhotoId_2 doesn't match, originPhotoId_2 [%s], resp %v",
			originPhotoId_2, resp))
	}

	if len(photo_1.PhotoId) == 0 || len(photo_1.PhotoUri) == 0 || photo_1.Likes != 0 {
		panic(fmt.Sprintf("received photo_1 is in wrong state, resp %v", resp))
	}

	photo_2 := resp.Photos[1]
	if photo_2.OriginPhotoId != originPhotoId {
		panic(fmt.Sprintf("received originPhotoId doesn't match, originPhotoId [%s], resp %v",
			originPhotoId, resp))
	}

	if len(photo_2.PhotoId) == 0 || len(photo_2.PhotoUri) == 0 || photo_2.Likes != 0 {
		panic(fmt.Sprintf("received photo_2 is in wrong state, resp %v", resp))
	}

	//now delete one photo
	respDel := apimodel.DeletePhoto(token_1, photo_2.PhotoId, true, lc)
	if len(respDel.ErrorCode) != 0 {
		panic(fmt.Sprintf("error when call /delete_photo, token [%s], photoId [%s], error code %s",
			token_1, photo_2.PhotoId, resp.ErrorCode))
	}

	resp = apimodel.GetOwnPhotos(token_1, apimodel.PhotoResolution480x640, true, lc)
	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error when call /get_own_photos, token [%s], resolution [%s], error code %s",
			token_1, apimodel.PhotoResolution480x640, resp.ErrorCode))
	}
	if len(resp.Photos) != 1 {
		panic(fmt.Sprintf("token_1 has wrong num of photos after delete one, result %v", resp))
	}

	//now check sorting order - last uploaded should be first
	photo_1 = resp.Photos[0]
	if photo_1.OriginPhotoId != originPhotoId_2 {
		panic(fmt.Sprintf("received originPhotoId_2 doesn't match, originPhotoId_2 [%s], resp %v",
			originPhotoId_2, resp))
	}

	if len(photo_1.PhotoId) == 0 || len(photo_1.PhotoUri) == 0 || photo_1.Likes != 0 {
		panic(fmt.Sprintf("received photo_1 is in wrong state, resp %v", resp))
	}
	apimodel.Anlogger.Infof(lc, "images.go : successfully complete image service complex test")
}

func UploadImage(token, sex string, baseNum, imageNum int, lc *lambdacontext.LambdaContext) string {
	getPresignResp := apimodel.GetPresignUrl(token, true, lc)
	if len(getPresignResp.ErrorCode) != 0 {
		panic(fmt.Sprintf("error call get presign, token [%s], base num [%d], imageNum [%d], error code %s",
			token, baseNum, imageNum, getPresignResp.ErrorCode))
	}
	image := apimodel.GenerateImage(sex == "male", fmt.Sprintf("%d.%d", baseNum, imageNum))
	apimodel.MakePutRequestWithContent(getPresignResp.Uri, image)
	return getPresignResp.OriginPhotoId
}
