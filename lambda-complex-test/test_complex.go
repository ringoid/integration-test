package main

import (
	"../apimodel"
	"../apitests"
	"context"
	basicLambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

func handler(ctx context.Context, request apimodel.CreateUsersRequest) error {
	lc, _ := lambdacontext.FromContext(ctx)
	apimodel.Anlogger.Infof(lc, "test_complex.go : complex test, req %v", request)

	apimodel.CleanAllDB(lc)

	apitests.LLMTest(lc)

	apimodel.CleanAllDB(lc)

	apimodel.Anlogger.Infof(lc, "test_complex.go : successfully complete complex test, req %v", request)
	return nil
}

func main() {
	basicLambda.Start(handler)
}
