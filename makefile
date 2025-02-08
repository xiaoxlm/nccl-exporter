.PHONY: build-amd-linux build-image

GO = GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go
LDFLAGS="-s -w"
GOBUILD=$(GO) build -a -ldflags $(LDFLAGS)

build-amd-linux:
	$(GOBUILD) -o bin/nccl-exporter-amd ./cmd/main.go

build-image:
	docker build --no-cache -f Dockerfile -t nccl-exporter:v1.0.0 .


