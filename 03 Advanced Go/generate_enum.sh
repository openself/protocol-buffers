#!/usr/bin/env bash
protoc -I src/ --go_out=src/ src/enum_example/enum_example.proto