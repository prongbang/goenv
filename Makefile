cover:
	go test -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o coverage.html