VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -s -w -X github.com/nednella/bootstrap.sh/internal.Version=$(VERSION)

build:
	cd cli && go build -trimpath -ldflags="$(LDFLAGS)" -o ../bin/bootstrap .

release:
	cd cli && GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 \
	  go build -trimpath -ldflags="$(LDFLAGS)" -o ../bin/bootstrap-darwin-arm64 .

run:
	cd cli && go run -ldflags="$(LDFLAGS)" . $(ARGS)

test:
	cd cli && go test ./...

fmt:
	cd cli && gofmt -w .

clean:
	rm -rf bin

.PHONY: build release run test fmt clean
