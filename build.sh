#!/bin/sh

go build -v -o bcWatch *.go
strip bcWatch

