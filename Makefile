#
# Makefile
# Paco Esteban, 2017-06-01 18:03
#

BINARY=duckdnsrefresh

VERSION=$(shell cat VERSION)
BUILD=`git rev-parse --short HEAD`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

.DEFAULT: build

build:
	mkdir -p ./build
	go build ${LDFLAGS} -o ./build/${BINARY}

xc_%:
	mkdir -p ./build
	$(eval SUFFIX=$(suffix $@))
	@echo "Cros compiling for: $(basename $*) $(SUFFIX:.%=%)"
	GOOS=$(basename $*) GOARCH=$(SUFFIX:.%=%) CGO_ENABLE=0 \
	     go build ${LDFLAGS} -o \
	     ./build/${BINARY}-${VERSION}-$(basename $*)-$(SUFFIX:.%=%)

buildall: xc_linux.386 xc_linux.amd64 xc_linux.arm\
	xc_openbsd.386 xc_openbsd.amd64 \
	xc_freebsd.386 xc_freebsd.amd64 \
	xc_darwin.amd64 \
	sign
sign:
	@echo "singing in the rain"
	cd ./build && \
		sha256sum * > SHA256 && \
		gpg --armor --detach-sign SHA256

clean:
	rm -rf ./build/*

.PHONY: clean build buildall sign

# vim:ft=make
#
