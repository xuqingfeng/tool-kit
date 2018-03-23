VERSION=`git describe --abbrev=0 --tags`

test:
	go vet ./... && go test -v $$(go list ./... | grep -v /vendor/)

fmt:
	go fmt ./...

build: fmt
	go build -ldflags "-w -s -X main.version=${VERSION}" -o out/wcl cmd/wcl/main.go
