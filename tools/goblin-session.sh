#!/bin/sh
# This is a test environment to see how usable goblin is on its own.

GOBLIN="$GOPATH/src/github.com/rkoesters/goblin"
export GOBLIN

PATH="$GOBLIN/bin"
export PATH

$SHELL
