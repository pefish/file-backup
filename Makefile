
DEFAULT: build-cur

ifeq ($(GOPATH),)
  GOPATH = $(HOME)/go
endif

build-cur:
	GOPATH=$(GOPATH) go install github.com/pefish/go-build-tool/cmd/...@latest
	$(GOPATH)/bin/go-build-tool

install: build-cur
	sudo install -C ./build/bin/linux/file-backup /usr/local/bin/file-backup

install-service: install
	sudo mkdir -p /etc/systemd/system
	sudo install -C -m 0644 ./script/file-backup.service /etc/systemd/system/file-backup.service
	sudo systemctl daemon-reload
	@echo
	@echo "file-backup service installed."

