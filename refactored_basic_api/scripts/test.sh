go test ./pkg/api/... -coverpkg=./pkg/api/... -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
xdg-open cover.html