generate:
	mockgen -source core/repository.go -destination core/repository_mock_test.go -package core
	mockgen -source core/util.go -destination core/util_mock_test.go -package core
test:
	go mod tidy
	ginkgo -r -v --trace --coverprofile=.cover-report.out ./...