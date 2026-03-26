#!/bin/bash



echo "Compilando windows....."
GOOS=windows GOARCH=amd64 go build -o ./builds/watchdog.exe main.go

echo "Compilando android....."
GOOS=android GOARCH=arm64 go build -o ./builds/watchdog-termux main.go 


echo "Compilando linux....."
go build -o ./builds/watchdog-linux main.go

