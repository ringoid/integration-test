test-all: clean test-deploy
stage-all: clean stage-deploy

build:
	@echo '--- Building generate-users-integration-test function ---'
	GOOS=linux go build lambda-create-users/create_users.go
	@echo '--- Building image-service-integration-test function ---'
	GOOS=linux go build lambda-image-test/test_image.go
	@echo '--- Building feeds-service-integration-test function ---'
	GOOS=linux go build lambda-feeds-test/test_feeds.go
	@echo '--- Building actions-service-integration-test function ---'
	GOOS=linux go build lambda-actions-test/test_actions.go
	@echo '--- Building auth-service-integration-test function ---'
	GOOS=linux go build lambda-auth-test/test_auth.go
	@echo '--- Building complex-integration-test function ---'
	GOOS=linux go build lambda-complex-test/test_complex.go

zip_lambda: build
	@echo '--- Zip generate-users-integration-test function ---'
	zip create_users.zip ./create_users
	@echo '--- Zip image-service-integration-test function ---'
	zip test_image.zip ./test_image
	@echo '--- Zip feeds-service-integration-test function ---'
	zip test_feeds.zip ./test_feeds
	@echo '--- Zip actions-service-integration-test function ---'
	zip test_actions.zip ./test_actions
	@echo '--- Zip auth-service-integration-test function ---'
	zip test_auth.zip ./test_auth
	@echo '--- Zip complex-integration-test function ---'
	zip test_complex.zip ./test_complex

test-deploy: zip_lambda
	@echo '--- Build lambda test ---'
	@echo 'Package template'
	sam package --template-file integration-test-template.yaml --s3-bucket ringoid-cloudformation-template --output-template-file integration-test-template-packaged.yaml
	@echo 'Deploy test-integration-test-stack'
	sam deploy --template-file integration-test-template-packaged.yaml --s3-bucket ringoid-cloudformation-template --stack-name test-integration-test-stack --capabilities CAPABILITY_IAM --parameter-overrides Env=test --no-fail-on-empty-changeset

stage-deploy: zip_lambda
	@echo '--- Build lambda stage ---'
	@echo 'Package template'
	sam package --template-file integration-test-template.yaml --s3-bucket ringoid-cloudformation-template --output-template-file integration-test-template-packaged.yaml
	@echo 'Deploy stage-integration-test-stack'
	sam deploy --template-file integration-test-template-packaged.yaml --s3-bucket ringoid-cloudformation-template --stack-name stage-integration-test-stack --capabilities CAPABILITY_IAM --parameter-overrides Env=stage --no-fail-on-empty-changeset

clean:
	@echo '--- Delete old artifacts ---'
	rm -rf create_users.zip
	rm -rf create_users
	rm -rf test_image.zip
	rm -rf test_image
	rm -rf test_feeds
	rm -rf test_feeds.zip
	rm -rf test_actions
	rm -rf test_actions.zip
	rm -rf test_auth.zip
	rm -rf test_auth
	rm -rf test_complex
	rm -rf test_complex.zip

