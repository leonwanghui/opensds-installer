.PHONY: all build package osdscf clean

all:build

build:osdscf

package:
	go get github.com/leonwanghui/opensds-installer/cmd/osdscf

osdscf:package
	mkdir -p  ./build/out/bin/
	go build -o ./build/out/bin/osdscf github.com/leonwanghui/opensds-installer/cmd/osdscf

clean:
	rm -rf ./build ./cmd/osdscf/osdscf
