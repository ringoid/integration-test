package main

import (
	"../apimodel"
	"../apitests"
	"context"
	basicLambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

func handler(ctx context.Context) error {
	lc, _ := lambdacontext.FromContext(ctx)
	apimodel.Anlogger.Infof(lc, "test_actions.go : start actions service integration test")

	apimodel.CleanAllDB(lc)

	apitests.ActionsWithOldBuildNum(lc)
	apitests.ActionsWithOldAccessToken(lc)
	apitests.ActionsWithEmptyActionType(lc)
	apitests.ActionsWithEmptySourceFeed(lc)
	apitests.ActionsWithEmptyTargetPhotoId(lc)
	apitests.ActionsWithEmptyTargetUserId(lc)
	apitests.ActionsWithNilActions(lc)
	apitests.ActionsWithWrongActionType(lc)
	apitests.ActionsWithWrongSourceFeed(lc)

	apimodel.CleanAllDB(lc)

	apimodel.Anlogger.Infof(lc, "test_actions.go : successfully end actions service integration test")
	return nil
}

func main() {
	basicLambda.Start(handler)
}
