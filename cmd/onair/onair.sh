#!/bin/sh
cd /opt/src/github.com/nh3000/broadcast/cmd/onair
go build onair.go
./onair -schedday MON -stationid WRRW -schedhour 23 -logging true