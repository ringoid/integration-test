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
	apimodel.Anlogger.Infof(lc, "test_feeds.go : start feeds service integration test")

	apimodel.CleanAllDB(lc)

	apitests.GetNewFacesWithOldBuildNum(lc)
	apitests.GetNewFacesWithOldAccessToken(lc)
	apitests.GetNewFacesWithWrongResolution(lc)

	apimodel.CleanAllDB(lc)

	apimodel.Anlogger.Infof(lc, "test_feeds.go : successfully end feeds service integration test")
	return nil
}

func main() {
	basicLambda.Start(handler)
}
