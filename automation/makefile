SERVER_BINARY_NAME := ismp-automation
AGENT_BINARY_NAME := ant-agent
go_build_cflags = -ldflags "-s -w"

.PHONY: build-web build-server rpm clean

build-web:
	@echo "Building web dist..."
	cd web
	yarn install
	yarn run build

build-server:
	@echo "Building server binary..."
	GOSUMDB=off CGO_ENABLED=0 GO111MODULE=on go build $(go_build_cflags) -o $(SERVER_BINARY_NAME) server/cmd/main.go

build-agent:
	@echo "Building agent binary..."
	GOSUMDB=off CGO_ENABLED=0 GO111MODULE=on go build $(go_build_cflags) -o $(AGENT_BINARY_NAME) agent/cmd/main.go

PWD_DIR := $(CURDIR)
PARENT_DIR := $(abspath $(PWD_DIR)/..)
ARCH    := $(shell uname -m)

rpm-server:
	@echo "Building tar..."
	rm -f $(PWD_DIR)/$(SERVER_BINARY_NAME).tar.gz
	cd $(PARENT_DIR) && tar -czvf $(SERVER_BINARY_NAME).tar.gz automation

	@echo "Moving tar..."
	mv $(PARENT_DIR)/$(SERVER_BINARY_NAME).tar.gz /root/rpmbuild/SOURCES/
	cp $(PWD_DIR)/server/scripts/ismp-automation.spec /root/rpmbuild/SPECS/

	@echo "Building rpm..."
	cd /root/rpmbuild/SPECS && rpmbuild -ba ismp-automation.spec

	@echo "Coping rpm..."
	cp /root/rpmbuild/RPMS/$(ARCH)/*.rpm $(PWD_DIR)/

rpm-agent:
	@echo "Building tar..."
	rm -f $(PWD_DIR)/$(AGENT_BINARY_NAME).tar.gz
	cd $(PARENT_DIR) && tar -czvf $(AGENT_BINARY_NAME).tar.gz automation

	@echo "Moving tar..."
	mv $(PARENT_DIR)/$(AGENT_BINARY_NAME).tar.gz /root/rpmbuild/SOURCES/
	cp $(PWD_DIR)/agent/scripts/ant-agent.spec /root/rpmbuild/SPECS/

	@echo "Building rpm..."
	cd /root/rpmbuild/SPECS && rpmbuild -ba ant-agent.spec

	@echo "Coping rpm..."
	cp /root/rpmbuild/RPMS/$(ARCH)/*.rpm $(PWD_DIR)/


clean:
	@echo "Cleaning up..."
	rm -rf web/dist
	rm -rf $(SERVER_BINARY_NAME)
	rm -rf $(AGENT_BINARY_NAME)