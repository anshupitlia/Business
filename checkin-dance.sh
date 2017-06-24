#!/bin/bash

git pull origin master -r
go test ./test/ && git push origin master

