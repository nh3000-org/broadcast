#!/bin/sh
cd /home/oem/go/src/github.com/nh3000-org/broadcast/cmd/exp
sudo -u postgres  pg_dump radio >  /run/media/oem/BackupWVOD/radio.sql
go build expcontent.go
./expcontent -exportpath /run/media/oem/BackupWVOD/
