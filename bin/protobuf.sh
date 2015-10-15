#!/bin/sh
dirName=${0%/*}
cd $dirName
cd ../library/types
protoc --go_out=./ ./*.proto