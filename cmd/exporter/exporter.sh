#!/bin/sh
cd /opt/src/github.com/nh3000/broadcast/cmd/exporter
go build exporter.go
./exporter -exportpath /media/oem/BackupWVOD/inventory
sudo -u postgres  pg_dump radio > sudo -u postgres  pg_dump radio > /media/oem/BackupWVOD/radio.sql.before