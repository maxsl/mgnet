#!/bin/bash
dirName=${0%/*}
cd $dirName
cd ../library/module/types
protoc --go_out=./ ./*.proto
