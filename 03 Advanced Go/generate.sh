#!/usr/bin/env bash
protoc -I src/ --go_out=src/ src/simple/simple.proto