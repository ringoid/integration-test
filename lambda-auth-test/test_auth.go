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
	apimodel.Anlogger.Infof(lc, "test_auth.go : start auth service integration test")

	apimodel.CleanAllDB(lc)

	apitests.CreateUserProfileWithOldBuildNum(lc)
	apitests.CreateUserProfileWithWrongYearOfBirth(lc)
	apitests.CreateUserProfileWithWrongSex(lc)

	apitests.UpdateSettingsWithOldBuildNum(lc)
	apitests.UpdateSettingsWithOldToken(lc)
	apitests.UpdateSettingsWithWrongDistance(lc)
	apitests.UpdateSettingsWithWrongPushLikes(lc)

	apimodel.CleanAllDB(lc)

	apimodel.Anlogger.Infof(lc, "test_auth.go : successfully end auth service integration test")
	return nil
}

func main() {
	basicLambda.Start(handler)
}
