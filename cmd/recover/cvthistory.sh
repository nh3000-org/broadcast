#!/bin/sh
cd /opt/src/github.com/nh3000/broadcast/cmd/recover
go build cvthistory.go
./cvthistory -rootimport /media/oem/BackupWVOD/radio.sql -stationid WVOD -t true -v true