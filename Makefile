BINARY  := distro-picker
VERSION := 0.1.0
LDFLAGS := -s -w -X main.version=$(VERSION)
OUTDIR  := bin

.PHONY: build run clean cross release

build:
	go build -ldflags="$(LDFLAGS)" -o $(OUTDIR)/$(BINARY) ./cmd/picker

run: build
	$(OUTDIR)/$(BINARY)

clean:
	rm -rf $(OUTDIR)

cross:
	GOOS=linux   GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o $(OUTDIR)/$(BINARY)-linux-amd64   ./cmd/picker
	GOOS=linux   GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o $(OUTDIR)/$(BINARY)-linux-arm64   ./cmd/picker
	GOOS=windows GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o $(OUTDIR)/$(BINARY)-windows-amd64.exe ./cmd/picker

release: cross
	@cd $(OUTDIR) && sha256sum $(BINARY)-* > checksums.txt
	@echo "Release artifacts in $(OUTDIR)/"

vet:
	go vet ./...

test:
	go test ./...
