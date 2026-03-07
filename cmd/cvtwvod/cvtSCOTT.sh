#!/bin/sh
cd /opt/src/github.com/nh3000/broadcast/cmd/onair
go build cvtscott.go
./cvtscott -rootimport "/media/oem/Backup Plus" -stationid WVOD 