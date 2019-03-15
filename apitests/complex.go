package apitests

import (
	"github.com/aws/aws-lambda-go/lambdacontext"
	"../apimodel"
	"github.com/ringoid/commons"
	"time"
	"fmt"
)

type Item struct {
	UserId  string
	PhotoId string
}

var resolution = apimodel.PhotoResolution1440x1920

//return sourceUserId, sourceToken, sp1, sp2, sp3 (photo ids of source user), result
func firstPhaseLikesYouTest(lc *lambdacontext.LambdaContext) (string, string, string, string, string, []Item) {
	source_token := CreateUser("male", lc)
	source_userId := apimodel.UserId(source_token, lc)

	originSp1 := UploadImage(source_token, "male", 1, 1, lc)
	time.Sleep(time.Second)
	originSp2 := UploadImage(source_token, "male", 1, 2, lc)
	time.Sleep(time.Second)
	originSp3 := UploadImage(source_token, "male", 1, 3, lc)
	time.Sleep(time.Second)
	sp1, _ := commons.GetResolutionPhotoId(source_userId, originSp1, resolution, apimodel.Anlogger, lc)
	sp2, _ := commons.GetResolutionPhotoId(source_userId, originSp2, resolution, apimodel.Anlogger, lc)
	sp3, _ := commons.GetResolutionPhotoId(source_userId, originSp3, resolution, apimodel.Anlogger, lc)

	target_1 := CreateUser("female", lc)
	target_1_userId := apimodel.UserId(target_1, lc)

	time.Sleep(time.Second)
	originT1p1 := UploadImage(target_1, "female", 100, 1, lc)
	time.Sleep(time.Second)
	originT1p2 := UploadImage(target_1, "female", 100, 2, lc)
	time.Sleep(time.Second)
	originT1p3 := UploadImage(target_1, "female", 100, 3, lc)
	time.Sleep(time.Second)
	originT1p4 := UploadImage(target_1, "female", 100, 4, lc)

	t1p1, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p1, resolution, apimodel.Anlogger, lc)
	t1p2, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p2, resolution, apimodel.Anlogger, lc)
	t1p3, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p3, resolution, apimodel.Anlogger, lc)
	t1p4, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p4, resolution, apimodel.Anlogger, lc)

	target_2 := CreateUser("female", lc)
	target_2_userId := apimodel.UserId(target_2, lc)
	time.Sleep(time.Second)
	originT2p1 := UploadImage(target_2, "female", 200, 1, lc)
	time.Sleep(time.Second)
	originT2p2 := UploadImage(target_2, "female", 200, 2, lc)
	time.Sleep(time.Second)
	originT2p3 := UploadImage(target_2, "female", 200, 3, lc)

	t2p1, _ := commons.GetResolutionPhotoId(target_2_userId, originT2p1, resolution, apimodel.Anlogger, lc)
	t2p2, _ := commons.GetResolutionPhotoId(target_2_userId, originT2p2, resolution, apimodel.Anlogger, lc)
	t2p3, _ := commons.GetResolutionPhotoId(target_2_userId, originT2p3, resolution, apimodel.Anlogger, lc)

	liker_1 := CreateUser("female", lc)
	liker_2 := CreateUser("female", lc)

	actionsForTarget_1 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   source_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   source_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp3,
			TargetUserId:   source_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_1, actionsForTarget_1, true, lc)

	actionsForTarget_2 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   source_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   source_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_2, actionsForTarget_2, true, lc)

	actionsForTarget_1 = []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   source_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   source_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_1, actionsForTarget_1, true, lc)

	actionsForTarget_2 = []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   source_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   source_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_2, actionsForTarget_2, true, lc)

	actionsForLiker_1 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p1,
			TargetUserId:   target_1_userId,
			ViewCount:      1,
			ViewTimeMillis: 10,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p2,
			TargetUserId:   target_1_userId,
			ViewCount:      2,
			ViewTimeMillis: 20,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p3,
			TargetUserId:   target_1_userId,
			ViewCount:      2,
			ViewTimeMillis: 30,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			ViewCount:      1,
			ViewTimeMillis: 10,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			ViewCount:      10,
			ViewTimeMillis: 10,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		//////-----------
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p1,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p2,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p3,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(liker_1, actionsForLiker_1, true, lc)

	actionsForLiker_2 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			ViewCount:      10,
			ViewTimeMillis: 100,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			ViewCount:      2,
			ViewTimeMillis: 100,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p3,
			TargetUserId:   target_2_userId,
			ViewCount:      10,
			ViewTimeMillis: 100,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		//////
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p3,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(liker_2, actionsForLiker_2, true, lc)

	actionsForSource := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(source_token, actionsForSource, true, lc)

	apimodel.GetUserSettings(source_token, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(target_1, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(target_2, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(liker_1, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(liker_2, true, lc)
	time.Sleep(time.Second)

	//Now test
	time.Sleep(2 * time.Second)
	resp := apimodel.GetLMM(source_token, resolution, 0, true, lc)
	apimodel.Anlogger.Debugf(lc, "firstPhaseLikesYouTest : source user id %s", source_userId)
	apimodel.Anlogger.Debugf(lc, "firstPhaseLikesYouTest : feed Resp : %v", resp)

	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("firstPhaseLikesYouTest : complex test, error code %s", resp.ErrorCode))
	}

	if resp.RepeatRequestAfter != 0 {
		panic(fmt.Sprintf("firstPhaseLikesYouTest : complex test, repeat request not zero, but %s", resp.RepeatRequestAfter))
	}

	photoReqiredOrder := []string{t2p1, t2p3, t2p2, t1p3, t1p2, t1p1, t1p4}

	//check that all profiles have unseen flag
	for _, each := range resp.LikesYou {
		if !each.Unseen {
			panic("firstPhaseLikesYouTest : complex test, new likes you have wrong unseen flag")
		}
	}

	photoActualOrder := make([]string, 0)
	for _, each := range resp.LikesYou {
		for _, eachP := range each.Photos {
			photoActualOrder = append(photoActualOrder, eachP.PhotoId)
		}
	}

	if len(photoReqiredOrder) != len(photoActualOrder) {
		panic("firstPhaseLikesYouTest : complex test, photoRequiredOrder.len != photoActualOrder.len")
	}

	for index, value := range photoReqiredOrder {
		if photoActualOrder[index] != value {
			apimodel.Anlogger.Debugf(lc, "complex.go : requiredOrder %v", photoReqiredOrder)
			apimodel.Anlogger.Debugf(lc, "complex.go : actualOrder   %v", photoActualOrder)
			panic("firstPhaseLikesYouTest : complex test, wrong order")
		}
	}

	resp = apimodel.GetLMM(source_token, resolution, time.Now().Round(time.Millisecond).UnixNano()/1000000, true, lc)
	apimodel.Anlogger.Debugf(lc, "firstPhaseLikesYouTest : second feed Resp : %v", resp)

	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("firstPhaseLikesYouTest : complex test, error code %s", resp.ErrorCode))
	}

	if resp.RepeatRequestAfter == 0 {
		panic(fmt.Sprintf("firstPhaseLikesYouTest : complex test, repeat request not zero, but %s", resp.RepeatRequestAfter))
	}

	result := []Item{
		Item{
			UserId:  target_2_userId,
			PhotoId: t2p1,
		},
		Item{
			UserId:  target_2_userId,
			PhotoId: t2p3,
		},
		Item{
			UserId:  target_2_userId,
			PhotoId: t2p2,
		},
		Item{
			UserId:  target_1_userId,
			PhotoId: t1p3,
		},
		Item{
			UserId:  target_1_userId,
			PhotoId: t1p2,
		},
		Item{
			UserId:  target_1_userId,
			PhotoId: t1p1,
		},
		Item{
			UserId:  target_1_userId,
			PhotoId: t1p4,
		},
	}
	return source_userId, source_token, sp1, sp2, sp3, result
}

func secondPhaseLikesYouTest(sourceUserId, sourceToke, sp1, sp2, sp3 string, firstPhaseResult []Item, lc *lambdacontext.LambdaContext) {
	actionsForSource := make([]apimodel.Action, 0)
	for _, item := range firstPhaseResult {
		actionsForSource = append(actionsForSource, apimodel.Action{
			SourceFeed:     apimodel.WhoLikedMeSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  item.PhotoId,
			TargetUserId:   item.UserId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		})
	}
	apimodel.Actions(sourceToke, actionsForSource, true, lc)

	target_1 := CreateUser("female", lc)
	target_1_userId := apimodel.UserId(target_1, lc)

	time.Sleep(time.Second)
	originT1p1 := UploadImage(target_1, "female", 100, 1, lc)
	time.Sleep(time.Second)
	originT1p2 := UploadImage(target_1, "female", 100, 2, lc)
	time.Sleep(time.Second)
	originT1p3 := UploadImage(target_1, "female", 100, 3, lc)
	time.Sleep(time.Second)
	originT1p4 := UploadImage(target_1, "female", 100, 4, lc)

	t1p1, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p1, resolution, apimodel.Anlogger, lc)
	t1p2, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p2, resolution, apimodel.Anlogger, lc)
	t1p3, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p3, resolution, apimodel.Anlogger, lc)
	t1p4, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p4, resolution, apimodel.Anlogger, lc)

	target_2 := CreateUser("female", lc)
	target_2_userId := apimodel.UserId(target_2, lc)
	time.Sleep(time.Second)
	originT2p1 := UploadImage(target_2, "female", 200, 1, lc)
	time.Sleep(time.Second)
	originT2p2 := UploadImage(target_2, "female", 200, 2, lc)
	time.Sleep(time.Second)
	originT2p3 := UploadImage(target_2, "female", 200, 3, lc)

	t2p1, _ := commons.GetResolutionPhotoId(target_2_userId, originT2p1, resolution, apimodel.Anlogger, lc)
	t2p2, _ := commons.GetResolutionPhotoId(target_2_userId, originT2p2, resolution, apimodel.Anlogger, lc)
	t2p3, _ := commons.GetResolutionPhotoId(target_2_userId, originT2p3, resolution, apimodel.Anlogger, lc)

	liker_1 := CreateUser("female", lc)
	liker_2 := CreateUser("female", lc)

	actionsForTarget_1 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   sourceUserId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   sourceUserId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp3,
			TargetUserId:   sourceUserId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_1, actionsForTarget_1, true, lc)

	actionsForTarget_2 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   sourceUserId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   sourceUserId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_2, actionsForTarget_2, true, lc)

	actionsForTarget_1 = []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   sourceUserId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   sourceUserId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_1, actionsForTarget_1, true, lc)

	actionsForTarget_2 = []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   sourceUserId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   sourceUserId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_2, actionsForTarget_2, true, lc)

	actionsForLiker_1 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p1,
			TargetUserId:   target_1_userId,
			ViewCount:      1,
			ViewTimeMillis: 10,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p2,
			TargetUserId:   target_1_userId,
			ViewCount:      1,
			ViewTimeMillis: 10,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p3,
			TargetUserId:   target_1_userId,
			ViewCount:      10,
			ViewTimeMillis: 100,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			ViewCount:      10,
			ViewTimeMillis: 10,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			ViewCount:      1,
			ViewTimeMillis: 100,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		///////
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p1,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p2,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p3,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(liker_1, actionsForLiker_1, true, lc)

	actionsForLiker_2 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			ViewCount:      10,
			ViewTimeMillis: 100,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			ViewCount:      10,
			ViewTimeMillis: 100,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p3,
			TargetUserId:   target_2_userId,
			ViewCount:      10,
			ViewTimeMillis: 100,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		/////
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p3,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(liker_2, actionsForLiker_2, true, lc)

	actionsForSource2 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(sourceToke, actionsForSource2, true, lc)

	apimodel.GetUserSettings(sourceToke, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(target_1, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(target_2, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(liker_1, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(liker_2, true, lc)
	time.Sleep(time.Second)

	//Now test
	time.Sleep(2 * time.Second)
	resp := apimodel.GetLMM(sourceToke, resolution, 0, true, lc)
	apimodel.Anlogger.Debugf(lc, "secondPhaseLikesYouTest : feed Resp : %v", resp)

	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("secondPhaseLikesYouTest : complex test, error code %s", resp.ErrorCode))
	}

	if resp.RepeatRequestAfter != 0 {
		panic(fmt.Sprintf("secondPhaseLikesYouTest : complex test, repeat request not zero, but %s", resp.RepeatRequestAfter))
	}

	photoReqiredOrderNewPart := []string{t2p1, t2p3, t2p2, t1p3, t1p2, t1p1, t1p4}
	apimodel.Anlogger.Debugf(lc, "secondPhaseLikesYouTest : photoReqiredOrderNewPart.len = %d", len(photoReqiredOrderNewPart))
	apimodel.Anlogger.Debugf(lc, "secondPhaseLikesYouTest : resp.LikesYou.len = %d", len(resp.LikesYou))

	if !resp.LikesYou[0].Unseen || !resp.LikesYou[1].Unseen {
		panic(fmt.Sprintf("secondPhaseLikesYouTest : wrong unseen flag in first part of likes"))
	}
	if resp.LikesYou[2].Unseen || resp.LikesYou[3].Unseen {
		panic(fmt.Sprintf("secondPhaseLikesYouTest : wrong unseen flag in second part of likes"))
	}

	photoActualOrderFullResult := make([]string, 0)
	for _, each := range resp.LikesYou {
		for _, eachP := range each.Photos {
			photoActualOrderFullResult = append(photoActualOrderFullResult, eachP.PhotoId)
		}
	}
	photoActualOrderNewPart := photoActualOrderFullResult[:len(photoReqiredOrderNewPart)]
	if len(photoReqiredOrderNewPart) != len(photoReqiredOrderNewPart) {
		panic("secondPhaseLikesYouTest : complex test, photoActualOrderNewPart.len != photoReqiredOrderNewPart.len")
	}

	for index, value := range photoReqiredOrderNewPart {
		if photoActualOrderNewPart[index] != value {
			apimodel.Anlogger.Errorf(lc, "secondPhaseLikesYouTest : required : %v", photoReqiredOrderNewPart)
			apimodel.Anlogger.Errorf(lc, "secondPhaseLikesYouTest : actual   : %v", photoActualOrderNewPart)
			panic("secondPhaseLikesYouTest : complex test, wrong order")
		}
	}

	//old part
	photoActualOrderOldPart := photoActualOrderFullResult[len(photoReqiredOrderNewPart):]
	if len(firstPhaseResult) != len(photoActualOrderOldPart) {
		panic("secondPhaseLikesYouTest : complex test, firstPhaseResult.len != photoActualOrderOldPart.len")
	}

	photoPredictedOrderOldPart := make([]string, 0)

	for _, item := range firstPhaseResult {
		photoPredictedOrderOldPart = append(photoPredictedOrderOldPart, item.PhotoId)
	}

	//тут играет то, что после того как я посмотрел фотки в LIKES_YOU
	//у них изменился порядок сортировки (условия смотрел их source теперь у всех = 1,
	//и решает теперь время загрузки фото.
	finalPhotoPredictOrderOldPart := make([]string, 0)
	finalPhotoPredictOrderOldPart = append(finalPhotoPredictOrderOldPart, photoPredictedOrderOldPart[2])
	finalPhotoPredictOrderOldPart = append(finalPhotoPredictOrderOldPart, photoPredictedOrderOldPart[0])
	finalPhotoPredictOrderOldPart = append(finalPhotoPredictOrderOldPart, photoPredictedOrderOldPart[1])
	finalPhotoPredictOrderOldPart = append(finalPhotoPredictOrderOldPart, photoPredictedOrderOldPart[3:]...)

	apimodel.Anlogger.Debugf(lc, "secondPhaseLikesYouTest : source userId %s", sourceUserId)
	apimodel.Anlogger.Debugf(lc, "secondPhaseLikesYouTest : predict order %v", photoPredictedOrderOldPart)
	apimodel.Anlogger.Debugf(lc, "secondPhaseLikesYouTest :  actual order %v", photoActualOrderOldPart)
	apimodel.Anlogger.Debugf(lc, "secondPhaseLikesYouTest :  final predict order %v", finalPhotoPredictOrderOldPart)

	for index, _ := range finalPhotoPredictOrderOldPart {
		if photoActualOrderOldPart[index] != finalPhotoPredictOrderOldPart[index] {
			panic("secondPhaseLikesYouTest : complex test, wrong order of old items")
		}
	}

	resp = apimodel.GetLMM(sourceToke, resolution, time.Now().Round(time.Millisecond).UnixNano()/1000000, true, lc)
	apimodel.Anlogger.Debugf(lc, "secondPhaseLikesYouTest : second feed Resp : %v", resp)

	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("secondPhaseLikesYouTest : complex test, error code %s", resp.ErrorCode))
	}

	if resp.RepeatRequestAfter == 0 {
		panic(fmt.Sprintf("secondPhaseLikesYouTest: complex test, repeat request not zero, but %s", resp.RepeatRequestAfter))
	}

}

func LLMTest(lc *lambdacontext.LambdaContext) {
	sourceUserId, sourceToke, sp1, sp2, sp3, firstPhaseResult := firstPhaseLikesYouTest(lc)
	secondPhaseLikesYouTest(sourceUserId, sourceToke, sp1, sp2, sp3, firstPhaseResult, lc)

	sourceUserId, sourceToke, sp1, sp2, sp3, firstPhaseResult = firstPhaseMatchesTest(lc)
	secondPhaseMatchesYouTest(sourceUserId, sourceToke, sp1, sp2, sp3, firstPhaseResult, lc)
}

//return sourceUserId, sourceToken, sp1, sp2, sp3 (photo ids of source user), result
func firstPhaseMatchesTest(lc *lambdacontext.LambdaContext) (string, string, string, string, string, []Item) {
	source_token := CreateUser("male", lc)
	source_userId := apimodel.UserId(source_token, lc)

	originSp1 := UploadImage(source_token, "male", 1, 1, lc)
	time.Sleep(time.Second)
	originSp2 := UploadImage(source_token, "male", 1, 2, lc)
	time.Sleep(time.Second)
	originSp3 := UploadImage(source_token, "male", 1, 3, lc)
	time.Sleep(time.Second)
	sp1, _ := commons.GetResolutionPhotoId(source_userId, originSp1, resolution, apimodel.Anlogger, lc)
	sp2, _ := commons.GetResolutionPhotoId(source_userId, originSp2, resolution, apimodel.Anlogger, lc)
	sp3, _ := commons.GetResolutionPhotoId(source_userId, originSp3, resolution, apimodel.Anlogger, lc)

	target_1 := CreateUser("female", lc)
	target_1_userId := apimodel.UserId(target_1, lc)

	time.Sleep(time.Second)
	originT1p1 := UploadImage(target_1, "female", 100, 1, lc)
	time.Sleep(time.Second)
	originT1p2 := UploadImage(target_1, "female", 100, 2, lc)
	time.Sleep(time.Second)
	originT1p3 := UploadImage(target_1, "female", 100, 3, lc)
	time.Sleep(time.Second)
	originT1p4 := UploadImage(target_1, "female", 100, 4, lc)

	t1p1, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p1, resolution, apimodel.Anlogger, lc)
	t1p2, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p2, resolution, apimodel.Anlogger, lc)
	t1p3, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p3, resolution, apimodel.Anlogger, lc)
	t1p4, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p4, resolution, apimodel.Anlogger, lc)

	target_2 := CreateUser("female", lc)
	target_2_userId := apimodel.UserId(target_2, lc)
	time.Sleep(time.Second)
	originT2p1 := UploadImage(target_2, "female", 200, 1, lc)
	time.Sleep(time.Second)
	originT2p2 := UploadImage(target_2, "female", 200, 2, lc)
	time.Sleep(time.Second)
	originT2p3 := UploadImage(target_2, "female", 200, 3, lc)

	t2p1, _ := commons.GetResolutionPhotoId(target_2_userId, originT2p1, resolution, apimodel.Anlogger, lc)
	t2p2, _ := commons.GetResolutionPhotoId(target_2_userId, originT2p2, resolution, apimodel.Anlogger, lc)
	t2p3, _ := commons.GetResolutionPhotoId(target_2_userId, originT2p3, resolution, apimodel.Anlogger, lc)

	liker_1 := CreateUser("female", lc)
	liker_2 := CreateUser("female", lc)

	actionsForSource := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p1,
			TargetUserId:   target_1_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p1,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},

		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p2,
			TargetUserId:   target_1_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p2,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},

		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},

		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(source_token, actionsForSource, true, lc)

	actionsForTarget_1 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   source_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   source_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp3,
			TargetUserId:   source_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_1, actionsForTarget_1, true, lc)

	actionsForTarget_2 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   source_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   source_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_2, actionsForTarget_2, true, lc)

	actionsForTarget_1 = []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   source_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   source_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_1, actionsForTarget_1, true, lc)

	actionsForTarget_2 = []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   source_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   source_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_2, actionsForTarget_2, true, lc)

	actionsForLiker_1 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p1,
			TargetUserId:   target_1_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p2,
			TargetUserId:   target_1_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p3,
			TargetUserId:   target_1_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},

		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p1,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p2,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p3,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(liker_1, actionsForLiker_1, true, lc)

	actionsForLiker_2 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p3,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},

		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p3,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(liker_2, actionsForLiker_2, true, lc)

	apimodel.GetUserSettings(source_token, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(target_1, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(target_2, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(liker_1, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(liker_2, true, lc)
	time.Sleep(time.Second)

	//Now test
	time.Sleep(2 * time.Second)
	resp := apimodel.GetLMM(source_token, resolution, 0, true, lc)
	apimodel.Anlogger.Debugf(lc, "firstPhaseMatchesTest : feed Resp : %v", resp)

	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("firstPhaseMatchesTest : complex test, error code %s", resp.ErrorCode))
	}

	if resp.RepeatRequestAfter != 0 {
		panic(fmt.Sprintf("firstPhaseMatchesTest : complex test, repeat request not zero, but %s", resp.RepeatRequestAfter))
	}

	photoReqiredOrder := []string{t1p3, t1p4, t1p2, t1p1, t2p3, t2p2, t2p1}

	for _, each := range resp.Matches {
		if !each.Unseen {
			panic(fmt.Sprintf("firstPhaseMatchesTest : complex test, wrong unseen flag for matches (first part)"))
		}
	}

	photoActualOrder := make([]string, 0)
	for _, each := range resp.Matches {
		for _, eachP := range each.Photos {
			photoActualOrder = append(photoActualOrder, eachP.PhotoId)
		}
	}

	if len(photoReqiredOrder) != len(photoActualOrder) {
		panic(fmt.Sprintf("firstPhaseMatchesTest : complex test, photoRequiredOrder.len != photoActualOrder.len (%d, %d)", len(photoReqiredOrder), len(photoActualOrder)))
	}

	for index, value := range photoReqiredOrder {
		if photoActualOrder[index] != value {
			panic("firstPhaseMatchesTest : complex test, wrong order")
		}
	}

	resp = apimodel.GetLMM(source_token, resolution, time.Now().Round(time.Millisecond).UnixNano()/1000000, true, lc)
	apimodel.Anlogger.Debugf(lc, "firstPhaseMatchesTest : second feed Resp : %v", resp)

	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("firstPhaseMatchesTest : complex test, error code %s", resp.ErrorCode))
	}

	if resp.RepeatRequestAfter == 0 {
		panic(fmt.Sprintf("firstPhaseTest : complex test, repeat request not zero, but %s", resp.RepeatRequestAfter))
	}

	result := []Item{
		Item{
			UserId:  target_1_userId,
			PhotoId: t1p3,
		},
		Item{
			UserId:  target_1_userId,
			PhotoId: t1p4,
		},
		Item{
			UserId:  target_1_userId,
			PhotoId: t1p2,
		},
		Item{
			UserId:  target_1_userId,
			PhotoId: t1p1,
		},
		Item{
			UserId:  target_2_userId,
			PhotoId: t2p3,
		},
		Item{
			UserId:  target_2_userId,
			PhotoId: t2p2,
		},
		Item{
			UserId:  target_2_userId,
			PhotoId: t2p1,
		},
	}
	return source_userId, source_token, sp1, sp2, sp3, result
}

func secondPhaseMatchesYouTest(sourceUserId, sourceToke, sp1, sp2, sp3 string, firstPhaseResult []Item, lc *lambdacontext.LambdaContext) {
	actionsForSource := make([]apimodel.Action, 0)
	for _, item := range firstPhaseResult {
		actionsForSource = append(actionsForSource, apimodel.Action{
			SourceFeed:     apimodel.MatchesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  item.PhotoId,
			TargetUserId:   item.UserId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		})
	}
	apimodel.Actions(sourceToke, actionsForSource, true, lc)

	target_1 := CreateUser("female", lc)
	target_1_userId := apimodel.UserId(target_1, lc)

	time.Sleep(time.Second)
	originT1p1 := UploadImage(target_1, "female", 100, 1, lc)
	time.Sleep(time.Second)
	originT1p2 := UploadImage(target_1, "female", 100, 2, lc)
	time.Sleep(time.Second)
	originT1p3 := UploadImage(target_1, "female", 100, 3, lc)
	time.Sleep(time.Second)
	originT1p4 := UploadImage(target_1, "female", 100, 4, lc)

	t1p1, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p1, resolution, apimodel.Anlogger, lc)
	t1p2, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p2, resolution, apimodel.Anlogger, lc)
	t1p3, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p3, resolution, apimodel.Anlogger, lc)
	t1p4, _ := commons.GetResolutionPhotoId(target_1_userId, originT1p4, resolution, apimodel.Anlogger, lc)

	target_2 := CreateUser("female", lc)
	target_2_userId := apimodel.UserId(target_2, lc)
	time.Sleep(time.Second)
	originT2p1 := UploadImage(target_2, "female", 200, 1, lc)
	time.Sleep(time.Second)
	originT2p2 := UploadImage(target_2, "female", 200, 2, lc)
	time.Sleep(time.Second)
	originT2p3 := UploadImage(target_2, "female", 200, 3, lc)

	t2p1, _ := commons.GetResolutionPhotoId(target_2_userId, originT2p1, resolution, apimodel.Anlogger, lc)
	t2p2, _ := commons.GetResolutionPhotoId(target_2_userId, originT2p2, resolution, apimodel.Anlogger, lc)
	t2p3, _ := commons.GetResolutionPhotoId(target_2_userId, originT2p3, resolution, apimodel.Anlogger, lc)

	liker_1 := CreateUser("female", lc)
	liker_2 := CreateUser("female", lc)

	actionsForSource = []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p1,
			TargetUserId:   target_1_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p1,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p2,
			TargetUserId:   target_1_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p2,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(sourceToke, actionsForSource, true, lc)

	actionsForTarget_1 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   sourceUserId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   sourceUserId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp3,
			TargetUserId:   sourceUserId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_1, actionsForTarget_1, true, lc)

	actionsForTarget_2 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   sourceUserId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   sourceUserId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_2, actionsForTarget_2, true, lc)

	actionsForTarget_1 = []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   sourceUserId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   sourceUserId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_1, actionsForTarget_1, true, lc)

	actionsForTarget_2 = []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp1,
			TargetUserId:   sourceUserId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  sp2,
			TargetUserId:   sourceUserId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(target_2, actionsForTarget_2, true, lc)

	actionsForLiker_1 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p1,
			TargetUserId:   target_1_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p2,
			TargetUserId:   target_1_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t1p3,
			TargetUserId:   target_1_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},

		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p1,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p2,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t1p3,
			TargetUserId:   target_1_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(liker_1, actionsForLiker_1, true, lc)

	actionsForLiker_2 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.ViewActionType,
			TargetPhotoId:  t2p3,
			TargetUserId:   target_2_userId,
			LikeCount:      0,
			ViewCount:      1,
			ViewTimeMillis: 2,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},

		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p1,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p2,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
		apimodel.Action{
			SourceFeed:     apimodel.NewFacesSourceFeed,
			ActionType:     apimodel.LikeActionType,
			TargetPhotoId:  t2p3,
			TargetUserId:   target_2_userId,
			LikeCount:      1,
			ViewCount:      0,
			ViewTimeMillis: 0,
			ActionTime:     time.Now().Round(time.Millisecond).UnixNano() / 1000000,
		},
	}
	apimodel.Actions(liker_2, actionsForLiker_2, true, lc)

	apimodel.GetUserSettings(sourceToke, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(target_1, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(target_2, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(liker_1, true, lc)
	time.Sleep(time.Second)
	apimodel.GetUserSettings(liker_2, true, lc)
	time.Sleep(time.Second)

	//Now test
	time.Sleep(2 * time.Second)
	resp := apimodel.GetLMM(sourceToke, resolution, 0, true, lc)
	apimodel.Anlogger.Debugf(lc, "secondPhaseMatchesYouTest : feed Resp : %v", resp)

	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("secondPhaseMatchesYouTest : complex test, error code %s", resp.ErrorCode))
	}

	if resp.RepeatRequestAfter != 0 {
		panic(fmt.Sprintf("secondPhaseMatchesYouTest : complex test, repeat request not zero, but %s", resp.RepeatRequestAfter))
	}

	photoReqiredOrderNewPart := []string{t1p3, t1p4, t1p2, t1p1, t2p3, t2p2, t2p1}

	if !resp.Matches[0].Unseen || !resp.Matches[1].Unseen {
		panic(fmt.Sprintf("secondPhaseMatchesYouTest : wrong unseen flag in first part of matches"))
	}
	if resp.Matches[2].Unseen || resp.Matches[3].Unseen {
		panic(fmt.Sprintf("secondPhaseMatchesYouTest : wrong unseen flag in second part of matches"))
	}

	photoActualFullOrder := make([]string, 0)
	for _, each := range resp.Matches {
		for _, eachP := range each.Photos {
			photoActualFullOrder = append(photoActualFullOrder, eachP.PhotoId)
		}
	}

	photoActualOrderNewPart := photoActualFullOrder[:len(photoReqiredOrderNewPart)]
	if len(photoReqiredOrderNewPart) != len(photoActualOrderNewPart) {
		panic("secondPhaseMatchesYouTest : complex test, photoActualOrderNewPart.len != photoReqiredOrderNewPart.len")
	}

	for index, value := range photoReqiredOrderNewPart {
		if photoActualOrderNewPart[index] != value {
			panic("secondPhaseMatchesYouTest : complex test, wrong order")
		}
	}

	//old part
	photoActualOrderOldPart := photoActualFullOrder[len(photoReqiredOrderNewPart):]

	photoPredictedOrderOldPart := make([]string, 0)

	for _, item := range firstPhaseResult {
		photoPredictedOrderOldPart = append(photoPredictedOrderOldPart, item.PhotoId)
	}

	//тут играет то, что после того как я посмотрел фотки в MATCHES_YOU
	//у них изменился порядок сортировки (условия смотрел их source теперь у конкурирующих = 1,
	//и решает теперь время сколько лайков у каждой.
	finalPhotoPredictOrderOldPart := make([]string, 0)
	finalPhotoPredictOrderOldPart = append(finalPhotoPredictOrderOldPart, photoPredictedOrderOldPart[0:4]...)
	finalPhotoPredictOrderOldPart = append(finalPhotoPredictOrderOldPart, photoPredictedOrderOldPart[5])
	finalPhotoPredictOrderOldPart = append(finalPhotoPredictOrderOldPart, photoPredictedOrderOldPart[4])
	finalPhotoPredictOrderOldPart = append(finalPhotoPredictOrderOldPart, photoPredictedOrderOldPart[6:]...)

	apimodel.Anlogger.Debugf(lc, "secondPhaseMatchesYouTest : source userId %s", sourceUserId)
	apimodel.Anlogger.Debugf(lc, "secondPhaseMatchesYouTest : predict order %v", photoPredictedOrderOldPart)
	apimodel.Anlogger.Debugf(lc, "secondPhaseMatchesYouTest :  actual order %v", photoActualOrderOldPart)
	apimodel.Anlogger.Debugf(lc, "secondPhaseMatchesYouTest :  final predict order %v", finalPhotoPredictOrderOldPart)

	if len(firstPhaseResult) != len(photoActualOrderOldPart) {
		panic(fmt.Sprintf("secondPhaseMatchesYouTest : complex test, firstPhaseResult.len != photoActualOrderOldPart.len, [%d] and [%d]", len(firstPhaseResult), len(photoActualOrderOldPart)))
	}

	for index, _ := range finalPhotoPredictOrderOldPart {
		if photoActualOrderOldPart[index] != finalPhotoPredictOrderOldPart[index] {
			panic("secondPhaseMatchesYouTest : complex test, wrong order of old items")
		}
	}

	resp = apimodel.GetLMM(sourceToke, resolution, time.Now().Round(time.Millisecond).UnixNano()/1000000, true, lc)
	apimodel.Anlogger.Debugf(lc, "secondPhaseMatchesYouTest : second feed Resp : %v", resp)

	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("secondPhaseMatchesYouTest : complex test, error code %s", resp.ErrorCode))
	}

	if resp.RepeatRequestAfter == 0 {
		panic(fmt.Sprintf("secondPhaseMatchesYouTest: complex test, repeat request not zero, but %s", resp.RepeatRequestAfter))
	}

}
