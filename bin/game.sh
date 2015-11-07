#!/bin/bash
dirName=${0%/*}
cd $dirName
cd ../game
go build -o ../bin/game
cd ../bin
./game