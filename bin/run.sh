#!/bin/bash

cf uninstall-plugin tree
go get ./...
cf install-plugin $GOPATH/bin/cf-plugin-tree
cf tree $@
