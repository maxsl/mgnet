#!/bin/bash
dirName=${0%/*}
cd $dirName
cd ../gate
go build -o ../bin/gate
cd ../bin
./gate