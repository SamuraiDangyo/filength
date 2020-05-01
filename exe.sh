#!/usr/bin/env bash

INSTALLDIR=/usr/bin

install() {
  go build filength.go
  sudo cp filength ${INSTALLDIR:?help}
}

run() {
  go run filength.go
}

if [ -z $1 ]; then
  install
else
  run
fi
