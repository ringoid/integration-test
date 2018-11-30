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
func firstPhaseTest(lc *lambdacontext.LambdaContext) (string, string, string, string, string, []Item) {
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
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.ViewActionType,
			TargetPhotoId: sp1,
			TargetUserId:  source_userId,
			LikeCount:     0,
			ViewCount:     1,
			ViewTimeSec:   2,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.ViewActionType,
			TargetPhotoId: sp2,
			TargetUserId:  source_userId,
			LikeCount:     0,
			ViewCount:     1,
			ViewTimeSec:   2,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.ViewActionType,
			TargetPhotoId: sp3,
			TargetUserId:  source_userId,
			LikeCount:     0,
			ViewCount:     1,
			ViewTimeSec:   2,
			ActionTime:    int(time.Now().Unix()),
		},
	}
	apimodel.Actions(target_1, actionsForTarget_1, true, lc)

	actionsForTarget_2 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.ViewActionType,
			TargetPhotoId: sp1,
			TargetUserId:  source_userId,
			LikeCount:     0,
			ViewCount:     1,
			ViewTimeSec:   2,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.ViewActionType,
			TargetPhotoId: sp2,
			TargetUserId:  source_userId,
			LikeCount:     0,
			ViewCount:     1,
			ViewTimeSec:   2,
			ActionTime:    int(time.Now().Unix()),
		},
	}
	apimodel.Actions(target_2, actionsForTarget_2, true, lc)

	actionsForTarget_1 = []apimodel.Action{
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: sp1,
			TargetUserId:  source_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: sp2,
			TargetUserId:  source_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
	}
	apimodel.Actions(target_1, actionsForTarget_1, true, lc)

	actionsForTarget_2 = []apimodel.Action{
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: sp1,
			TargetUserId:  source_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: sp2,
			TargetUserId:  source_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
	}
	apimodel.Actions(target_2, actionsForTarget_2, true, lc)

	actionsForLiker_1 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t1p1,
			TargetUserId:  target_1_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t1p2,
			TargetUserId:  target_1_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t1p3,
			TargetUserId:  target_1_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t2p1,
			TargetUserId:  target_2_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t2p2,
			TargetUserId:  target_2_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
	}
	apimodel.Actions(liker_1, actionsForLiker_1, true, lc)

	actionsForLiker_2 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t2p1,
			TargetUserId:  target_2_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t2p2,
			TargetUserId:  target_2_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t2p3,
			TargetUserId:  target_2_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
	}
	apimodel.Actions(liker_2, actionsForLiker_2, true, lc)

	actionsForSource := []apimodel.Action{
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.ViewActionType,
			TargetPhotoId: t2p2,
			TargetUserId:  target_2_userId,
			LikeCount:     0,
			ViewCount:     1,
			ViewTimeSec:   2,
			ActionTime:    int(time.Now().Unix()),
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
	apimodel.Anlogger.Debugf(lc, "firstPhaseTest : feed Resp : %v", resp)

	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("firstPhaseTest : complex test, error code %s", resp.ErrorCode))
	}

	if resp.RepeatRequestAfterSec != 0 {
		panic(fmt.Sprintf("firstPhaseTest : complex test, repeat request not zero, but %s", resp.RepeatRequestAfterSec))
	}

	photoReqiredOrder := []string{t2p1, t2p3, t2p2, t1p3, t1p2, t1p1, t1p4}

	photoActualOrder := make([]string, 0)
	for _, each := range resp.LikesYouNewProfiles {
		for _, eachP := range each.Photos {
			photoActualOrder = append(photoActualOrder, eachP.PhotoId)
		}
	}

	if len(photoReqiredOrder) != len(photoActualOrder) {
		panic("firstPhaseTest : complex test, photoRequiredOrder.len != photoActualOrder.len")
	}

	for index, value := range photoReqiredOrder {
		if photoActualOrder[index] != value {
			panic("firstPhaseTest : complex test, wrong order")
		}
	}

	resp = apimodel.GetLMM(source_token, resolution, int(time.Now().Unix()), true, lc)
	apimodel.Anlogger.Debugf(lc, "firstPhaseTest : second feed Resp : %v", resp)

	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("firstPhaseTest : complex test, error code %s", resp.ErrorCode))
	}

	if resp.RepeatRequestAfterSec == 0 {
		panic(fmt.Sprintf("firstPhaseTest : complex test, repeat request not zero, but %s", resp.RepeatRequestAfterSec))
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

func secondPhaseTest(sourceUserId, sourceToke, sp1, sp2, sp3 string, firstPhaseResult []Item, lc *lambdacontext.LambdaContext) {
	actionsForSource := make([]apimodel.Action, 0)
	for _, item := range firstPhaseResult {
		actionsForSource = append(actionsForSource, apimodel.Action{
			SourceFeed:    apimodel.WhoLikedMeSourceFeed,
			ActionType:    apimodel.ViewActionType,
			TargetPhotoId: item.PhotoId,
			TargetUserId:  item.UserId,
			LikeCount:     0,
			ViewCount:     1,
			ViewTimeSec:   2,
			ActionTime:    int(time.Now().Unix()),
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
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.ViewActionType,
			TargetPhotoId: sp1,
			TargetUserId:  sourceUserId,
			LikeCount:     0,
			ViewCount:     1,
			ViewTimeSec:   2,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.ViewActionType,
			TargetPhotoId: sp2,
			TargetUserId:  sourceUserId,
			LikeCount:     0,
			ViewCount:     1,
			ViewTimeSec:   2,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.ViewActionType,
			TargetPhotoId: sp3,
			TargetUserId:  sourceUserId,
			LikeCount:     0,
			ViewCount:     1,
			ViewTimeSec:   2,
			ActionTime:    int(time.Now().Unix()),
		},
	}
	apimodel.Actions(target_1, actionsForTarget_1, true, lc)

	actionsForTarget_2 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.ViewActionType,
			TargetPhotoId: sp1,
			TargetUserId:  sourceUserId,
			LikeCount:     0,
			ViewCount:     1,
			ViewTimeSec:   2,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.ViewActionType,
			TargetPhotoId: sp2,
			TargetUserId:  sourceUserId,
			LikeCount:     0,
			ViewCount:     1,
			ViewTimeSec:   2,
			ActionTime:    int(time.Now().Unix()),
		},
	}
	apimodel.Actions(target_2, actionsForTarget_2, true, lc)

	actionsForTarget_1 = []apimodel.Action{
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: sp1,
			TargetUserId:  sourceUserId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: sp2,
			TargetUserId:  sourceUserId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
	}
	apimodel.Actions(target_1, actionsForTarget_1, true, lc)

	actionsForTarget_2 = []apimodel.Action{
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: sp1,
			TargetUserId:  sourceUserId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: sp2,
			TargetUserId:  sourceUserId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
	}
	apimodel.Actions(target_2, actionsForTarget_2, true, lc)

	actionsForLiker_1 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t1p1,
			TargetUserId:  target_1_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t1p2,
			TargetUserId:  target_1_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t1p3,
			TargetUserId:  target_1_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t2p1,
			TargetUserId:  target_2_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t2p2,
			TargetUserId:  target_2_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
	}
	apimodel.Actions(liker_1, actionsForLiker_1, true, lc)

	actionsForLiker_2 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t2p1,
			TargetUserId:  target_2_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t2p2,
			TargetUserId:  target_2_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.LikeActionType,
			TargetPhotoId: t2p3,
			TargetUserId:  target_2_userId,
			LikeCount:     1,
			ViewCount:     0,
			ViewTimeSec:   0,
			ActionTime:    int(time.Now().Unix()),
		},
	}
	apimodel.Actions(liker_2, actionsForLiker_2, true, lc)

	actionsForSource2 := []apimodel.Action{
		apimodel.Action{
			SourceFeed:    apimodel.NewFacesSourceFeed,
			ActionType:    apimodel.ViewActionType,
			TargetPhotoId: t2p2,
			TargetUserId:  target_2_userId,
			LikeCount:     0,
			ViewCount:     1,
			ViewTimeSec:   2,
			ActionTime:    int(time.Now().Unix()),
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
	apimodel.Anlogger.Debugf(lc, "secondPhaseTest : feed Resp : %v", resp)

	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("secondPhaseTest : complex test, error code %s", resp.ErrorCode))
	}

	if resp.RepeatRequestAfterSec != 0 {
		panic(fmt.Sprintf("secondPhaseTest : complex test, repeat request not zero, but %s", resp.RepeatRequestAfterSec))
	}

	photoReqiredOrderNewPart := []string{t2p1, t2p3, t2p2, t1p3, t1p2, t1p1, t1p4}

	photoActualOrderNewPart := make([]string, 0)
	for _, each := range resp.LikesYouNewProfiles {
		for _, eachP := range each.Photos {
			photoActualOrderNewPart = append(photoActualOrderNewPart, eachP.PhotoId)
		}
	}

	if len(photoReqiredOrderNewPart) != len(photoActualOrderNewPart) {
		panic("secondPhaseTest : complex test, photoActualOrderNewPart.len != photoReqiredOrderNewPart.len")
	}

	for index, value := range photoReqiredOrderNewPart {
		if photoActualOrderNewPart[index] != value {
			panic("secondPhaseTest : complex test, wrong order")
		}
	}

	//old part
	photoActualOrderOldPart := make([]string, 0)
	for _, each := range resp.LikesYouOldProfiles {
		for _, eachP := range each.Photos {
			photoActualOrderOldPart = append(photoActualOrderOldPart, eachP.PhotoId)
		}
	}

	if len(firstPhaseResult) != len(photoActualOrderOldPart) {
		panic("secondPhaseTest : complex test, firstPhaseResult.len != photoActualOrderOldPart.len")
	}

	for index, item := range firstPhaseResult {
		if photoActualOrderOldPart[index] != item.PhotoId {
			panic("secondPhaseTest : complex test, wrong orde of old items")
		}
	}

	resp = apimodel.GetLMM(sourceToke, resolution, int(time.Now().Unix()), true, lc)
	apimodel.Anlogger.Debugf(lc, "secondPhaseTest : second feed Resp : %v", resp)

	if len(resp.ErrorCode) != 0 {
		panic(fmt.Sprintf("secondPhaseTest : complex test, error code %s", resp.ErrorCode))
	}

	if resp.RepeatRequestAfterSec == 0 {
		panic(fmt.Sprintf("secondPhaseTest: complex test, repeat request not zero, but %s", resp.RepeatRequestAfterSec))
	}

}

func LLMTest(lc *lambdacontext.LambdaContext) {
	sourceUserId, sourceToke, sp1, sp2, sp3, firstPhaseResult := firstPhaseTest(lc)
	secondPhaseTest(sourceUserId, sourceToke, sp1, sp2, sp3, firstPhaseResult, lc)
}
