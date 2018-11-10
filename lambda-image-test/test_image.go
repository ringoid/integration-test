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
	apimodel.Anlogger.Infof(lc, "test_image.go : start image service integration test")

	apimodel.CleanAllDB(lc)

	apitests.GetPresignWithOldBuildNumTest(lc)
	apitests.GetPresignWithOldAccessTokenTest(lc)
	apitests.GetOwnPhotosWithWrongResolution(lc)

	apitests.GetOwnPhotosWithOldBuildNum(lc)
	apitests.GetOwnPhotosWithOldToken(lc)
	apitests.GetOwnPhotosWithWrongResolution(lc)

	apitests.DeletePhotoWithOldBuildNum(lc)
	apitests.DeletePhotoWithOldAccessToken(lc)
	apitests.DeletePhotoWithEmptyPhotoId(lc)

	apitests.ImageTest(lc)

	apimodel.CleanAllDB(lc)

	apimodel.Anlogger.Infof(lc, "test_image.go : successfully end image service integration test")
	return nil
}

func main() {
	basicLambda.Start(handler)
}
