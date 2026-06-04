VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -s -w -X github.com/nednella/bootstrap.sh/internal.Version=$(VERSION)
LOCAL_INSTALL_DEST := /usr/local/bin/bootstrap

build:
	cd cli && go build -trimpath -ldflags="$(LDFLAGS)" -o ../bin/bootstrap .

install-dev:
	cd cli && go build -o ../bin/bootstrap .
	sudo mv bin/bootstrap $(LOCAL_INSTALL_DEST)

uninstall-dev:
	sudo rm -f $(LOCAL_INSTALL_DEST)

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

.PHONY: build install-dev uninstall-dev release run test fmt clean
