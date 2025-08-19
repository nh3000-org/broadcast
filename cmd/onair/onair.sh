#!/bin/sh
cd /opt/src/github.com/nh3000/broadcast/cmd/onair
go build onair.go
./onair -schedday TUE -stationid WVOD -schedhour 06 -logging true
