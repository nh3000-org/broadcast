#!/bin/sh
cd /opt/src/github.com/nh3000/broadcast/cmd/cvtwvod
go build cvtSCOTT.go
./cvtSCOTT -rootimport "/run/media/oem/Backup Plus" -stationid WVOD  -test false
