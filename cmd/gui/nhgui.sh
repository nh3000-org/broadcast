#!/bin/sh
cd /opt/src/github.com/nh3000/broadcast/cmd/gui
export FYNE_SCALE='2'
go build nhgui.go
./nhgui 
