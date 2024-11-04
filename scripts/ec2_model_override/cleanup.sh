#!/bin/bash

rm -rf ./vendor
go mod edit -dropreplace github.com/nholuongut/nholuongut-sdk-go
go mod tidy
