#!/usr/bin/env bash

command -v upx >/dev/null 2>&1 || { echo >&2 "Could not find required 'upx' program. Aborting."; exit 1; }
command -v go >/dev/null 2>&1  || { echo >&2 "Could not find required 'go' program. Aborting."; exit 1; }

DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

VERSION_SHORT=$(cat ${DIR}/VERSION)
VERSION_LONG=$(echo "v$(cat ${DIR}/VERSION) ($(git describe --abbrev=8 --always --long))")
cat ${DIR}/logo.txt | sed -e "s/VERSION/${VERSION_LONG}/g" > ${DIR}/../src/cmd/logo.txt

SRCPATH=${DIR}/..

DSTPATH=$1
if [ -z $DSTPATH ]; then
  DSTPATH="${DIR}/../distr"
fi

function cleanup {
  rm -rf ${DSTPATH}/bin || true
}

function build() {
  mkdir -p ${DSTPATH}/bin/$6/$7/
  cd ${SRCPATH}/src
  echo "> $6.$4"
  RES=$(GOOS=$3 GOARCH=$4 GOARM=$5 go build -buildvcs=false -ldflags "-s -w" -o ${DSTPATH}/bin/$6/$7/$2$8 .)
  if [[ "linux" == "$3" && $RES ]]; then
    mv ${DSTPATH}/bin/$6/$7/$2$8 ${DSTPATH}/bin/$6/$7/$2$8.bin
    RES=$(upx -9 -o ${DSTPATH}/bin/$6/$7/$2$8 ${DSTPATH}/bin/$6/$7/$2$8.bin 1>/dev/null)
    rm ${DSTPATH}/bin/$6/$7/$2$8.bin
  fi
  return $RES
}

echo                                                                && \
echo "Building keito"                                               && \
echo                                                                && \
cleanup                                                             && \
cd src && go mod tidy                                               && \
build src keito darwin  amd64   ""    macos   x86_64  ""            && \
build src keito darwin  arm64   ""    macos   aarch64 ""            && \
build src keito linux   amd64   ""    linux   amd64   ""            && \
build src keito linux   arm     7     linux   armhf   ""            && \
build src keito linux   arm64   7     linux   armhf64 ""            && \
build src keito windows amd64   ""    win     x86_64  ".exe"        && \
build src keito windows 386     ""    win     x86     ".exe"        && \
echo                                                                && \
echo "Done!"                                                        && \
echo


# --- cut here -----------------------------------------------------------------------------------------

# go tool dist list

# GOOS
#Linux 	        linux
#MacOS X 	      darwin
#Windows 	      windows
#FreeBSD 	      freebsd
#NetBSD 	      netbsd
#OpenBSD 	      openbsd
#DragonFly BSD 	dragonfly
#Plan 9 	      plan9
#Native Client 	nacl
#Android 	      android

# GOARCH
#x386 	                  386
#AMD64 	                  amd64
#AMD64 с 32-указателями 	amd64p32
#ARM 	                    arm
#ARM 	                    arm64

# GOARM
# armel (softfloat)               GOARM=5
# armhf (hardware floating point) GOARM=6 / GOARM=7

