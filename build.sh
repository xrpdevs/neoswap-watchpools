#!/bin/sh

git pull 
go build -v -o bcWatch *.go
strip bcWatch

