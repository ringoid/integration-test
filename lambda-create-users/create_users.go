package main

import (
	"../apimodel"
	"context"
	basicLambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"fmt"
	"math/rand"
)

func handler(ctx context.Context, request apimodel.CreateUsersRequest) error {
	lc, _ := lambdacontext.FromContext(ctx)
	apimodel.Anlogger.Infof(lc, "create_users.go : create users, req %v", request)
	for i := 0; i < request.MenNum; i++ {
		generateUser(request.NextNum, i, "male", lc)
	}

	for i := 0; i < request.WomenNum; i++ {
		generateUser(request.NextNum, i, "female", lc)
	}

	apimodel.Anlogger.Infof(lc, "create_users.go : successfully create users, req %v", request)
	return nil
}

func generateUser(baseNum, i int, sex string, lc *lambdacontext.LambdaContext) {
	authResp := apimodel.Start(7, fmt.Sprintf("%d%d", baseNum, i), lc)
	verifyResp := apimodel.Complete(authResp.SessionId, "fake_code", lc)
	apimodel.CreateUserProfile(verifyResp.AccessToken, 1980+rand.Intn(20), sex, true, lc)
	for j := 0; j < 3+rand.Intn(3); j++ {
		image := apimodel.GenerateImage("male" == sex, fmt.Sprintf("%d.%d.%d", baseNum,	 i, j))
		getPresignResp := apimodel.GetPresignUrl(verifyResp.AccessToken, true, lc)
		apimodel.MakePutRequestWithContent(getPresignResp.Uri, image)
	}
}

func main() {
	basicLambda.Start(handler)
}
