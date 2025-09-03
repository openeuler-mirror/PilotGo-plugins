BINARY_NAME := PilotGo-plugin-automation
go_build_cflags = -ldflags "-s -w"

.PHONY: build-web build-server rpm clean

build-web:
	@echo "Building web dist..."
	cd web
	yarn install
	yarn run build

build-server:
	@echo "Building server binary..."
	GOSUMDB=off GO111MODULE=on go build $(go_build_cflags) -o $(BINARY_NAME) server/cmd/main.go

PWD_DIR := $(CURDIR)
PARENT_DIR := $(abspath $(PWD_DIR)/..)
ARCH    := $(shell uname -m)

rpm:
	@echo "Building tar..."
	rm -f $(PWD_DIR)/$(BINARY_NAME).tar.gz
	cd $(PARENT_DIR) && tar -czvf $(BINARY_NAME).tar.gz automation

	@echo "Moving tar..."
	mv $(PARENT_DIR)/$(BINARY_NAME).tar.gz /root/rpmbuild/SOURCES/
	cp $(PWD_DIR)/server/scripts/PilotGo-plugin-automation.spec /root/rpmbuild/SPECS/

	@echo "Building rpm..."
	cd /root/rpmbuild/SPECS && rpmbuild -ba PilotGo-plugin-automation.spec

	@echo "Coping rpm..."
	cp /root/rpmbuild/RPMS/$(ARCH)/*.rpm $(PWD_DIR)/


clean:
	@echo "Cleaning up..."
	rm -rf web/dist
	rm -rf $(BINARY_NAME)