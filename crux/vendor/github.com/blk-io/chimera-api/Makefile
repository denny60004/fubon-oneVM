# The import path is where your repository can be found.
# To import subpackages, always prepend the full import path.
# If you change this, run `make clean`. Read more: https://git.io/vM7zV
IMPORT_PATH := github.com/blk-io/chimera-api
Q := $(if $V,,@)
export GOPATH := $(CURDIR)/.GOPATH

protofiles: protos grpc

protos:
	$Q protoc -I proto/ \
		-I .GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:chimera \
		proto/*.proto

grpc:
	$Q protoc -I proto/ \
		-I .GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:chimera \
		proto/grpc.proto

.PHONY: build
build: protofiles


.PHONY: setup
setup: clean .GOPATH/.ok
	@if ! grep "/.GOPATH" .gitignore > /dev/null 2>&1; then \
	    echo "/.GOPATH" >> .gitignore; \
	    echo "/bin" >> .gitignore; \
	fi
	$Q curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
	$Q unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
	$Q mv protoc3/bin/* .GOPATH/bin/
	$Q mv protoc3/include/* .GOPATH/include/
	$Q rm -r protoc3
	$Q rm protoc-3.5.1-linux-x86_64.zip
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway


.GOPATH/.ok:
	$Q echo $(IMPORT_PATH)
	$Q mkdir -p "$(dir .GOPATH/src/$(IMPORT_PATH))"
	$Q ln -s ../../../.. ".GOPATH/src/$(IMPORT_PATH)"
	$Q mkdir -p .GOPATH/test .GOPATH/cover
	$Q mkdir -p bin
	$Q ln -s ../bin .GOPATH/bin
	$Q touch $@

clean:
	$Q rm -rf bin .GOPATH
	$Q rm -f ./chimera/*