package apitests

import (
	"github.com/aws/aws-lambda-go/lambdacontext"
	"../apimodel"
	"github.com/ringoid/commons"
	"time"
	"fmt"
)

const (
	newFacesDefaultLimit = 5
	iteratinosNum        = 100
)

//return sourceUserId, sourceToken, sp1, sp2, sp3 (photo ids of source user), result
func FirstSpeedLikesYouTest(lc *lambdacontext.LambdaContext) {

	//source user creation
	source_token := CreateUser("male", lc)
	sourceUserId := apimodel.UserId(source_token, lc)

	UploadImage(source_token, "male", 1, 1, lc)
	time.Sleep(100 * time.Millisecond)
	UploadImage(source_token, "male", 1, 2, lc)
	time.Sleep(100 * time.Millisecond)
	UploadImage(source_token, "male", 1, 3, lc)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("(main info test) Finish source users creation, userId [%s], at %v\n", sourceUserId, time.Now())
	//prepare map for target users
	targetUsers := make(map[string]string)
	fmt.Printf("(main info test) Start target users creation at %v\n", time.Now())
	for i := 0; i < iteratinosNum; i++ {
		targetToken := CreateUser("female", lc)
		targetUserId := apimodel.UserId(targetToken, lc)
		originImageId := UploadImage(targetToken, "female", i, 1, lc)
		resolutionImageId, _ := commons.GetResolutionPhotoId(targetUserId, originImageId, resolution, apimodel.Anlogger, lc)

		targetUsers[fmt.Sprintf("targetToken_%d", i)] = targetToken
		targetUsers[fmt.Sprintf("targetUserId_%d", i)] = targetUserId
		targetUsers[fmt.Sprintf("resoulutionImageId_%d", i)] = resolutionImageId
	}
	fmt.Printf("(main info test) Finish target users creation at %v\n", time.Now())
	time.Sleep(time.Second * 5)

	var lastActionTime int64
	for i := 0; i < iteratinosNum/newFacesDefaultLimit; i++ {
		var resp apimodel.GetNewFacesResp

		for {
			fmt.Printf("(main info test) Iteration [%d], start asking at %v\n", i, time.Now())
			resp = apimodel.GetNewFaces(source_token, resolution, lastActionTime, true, lc)
			fmt.Printf("(main info test) Iteration [%d], receive response at %v\n", i, time.Now())

			if len(resp.ErrorCode) != 0 {
				fmt.Printf("(main info test) Iteration [%d], error while fetch NewFaces : %s\n", i, resp.ErrorCode)
				panic(fmt.Sprintf("(main info test) Iteration [%d], error while fetch NewFaces : %s\n", i, resp.ErrorCode))
			}
			if len(resp.Profiles) != newFacesDefaultLimit && resp.RepeatRequestAfterSec == 0 {
				fmt.Printf("(main info test) Iteration [%d], error while fetch NewFaces : result profiles num [%d]\n", i, 0)
				panic(fmt.Sprintf("(main info test) Iteration [%d], error while fetch NewFaces : result profiles num [%d]\n", i, 0))
			}

			if resp.RepeatRequestAfterSec != 0 {
				fmt.Printf("(main info test) Iteration [%d], repeat request after sec [%d]\n", i, resp.RepeatRequestAfterSec)
				time.Sleep(100 * time.Millisecond)
			} else {
				fmt.Printf("(main info test) Iteration [%d], receive response with profiles [%d] at %v\n", i, len(resp.Profiles), time.Now())
				break
			}
		}

		actions := make([]apimodel.Action, 0)
		for _, eachProfile := range resp.Profiles {
			for _, eachPhoto := range eachProfile.Photos {
				actions = append(actions, apimodel.Action{
					SourceFeed:     apimodel.NewFacesSourceFeed,
					ActionType:     apimodel.ViewActionType,
					TargetPhotoId:  eachPhoto.PhotoId,
					TargetUserId:   eachProfile.UserId,
					LikeCount:      0,
					ViewCount:      1,
					ViewTimeMillis: 2,
					ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
				})

				lastActionTime = time.Now().Round(time.Millisecond).UnixNano() / 1000000
				actions = append(actions, apimodel.Action{
					SourceFeed:     apimodel.NewFacesSourceFeed,
					ActionType:     apimodel.LikeActionType,
					TargetPhotoId:  eachPhoto.PhotoId,
					TargetUserId:   eachProfile.UserId,
					LikeCount:      1,
					ViewCount:      0,
					ViewTimeMillis: 0,
					ActionTime:     lastActionTime,
				})
			}
		}

		fmt.Printf("(main info test) Iteration [%d], start sending [%d] nums actions at %v\n", i, len(actions), time.Now())
		apimodel.Actions(source_token, actions, true, lc)
		fmt.Printf("(main info test) Iteration [%d], end sending [%d] nums actions at %v\n", i, len(actions), time.Now())
	}
}
