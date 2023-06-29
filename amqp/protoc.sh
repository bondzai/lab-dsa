#!/bin/sh

protoc --python_out=. task.proto
protoc --go_out=. task.proto