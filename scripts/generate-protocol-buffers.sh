#!/bin/sh

protoc --proto_path=. --proto_path=include lib/global-bus.proto --go_out=../../..
protoc --proto_path=. --proto_path=include samples/my-events.proto --go_out=.
