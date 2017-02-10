all: lint wind

GCFLAGS = -l -N
LDFLAGS += -X "github.com/godbs/util.BuildTime=$(shell date +'%F %T %z')"
LDFLAGS += -X "github.com/godbs/util.BuildVersion=$(shell git rev-parse HEAD)"

FILES := $$(find . -name '*.go' | grep -vE 'vendor') 
SOURCE_PATH := server

golint:
	go get github.com/golang/lint/golint  

godep:
	go get github.com/tools/godep

lint: golint
	@for path in $(SOURCE_PATH); do echo "golint $$path"; golint $$path"/..."; done;
	@for path in $(SOURCE_PATH); do echo "gofmt -s -l -w $$path";  gofmt -s -l -w $$path;  done;
	go tool vet $(FILES) 2>&1
	go tool vet --shadow $(FILES) 2>&1

clean:
	@rm -rf bin


wind:godep
	go build -o bin/$@ -gcflags="$(GCFLAGS)" -ldflags '$(LDFLAGS)' ./main.go

test:
	@for path in $(SOURCE_PATH); do echo "go test ./$$path"; go test "./"$$path; done;


