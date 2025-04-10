#EX: tail -f log.file | nhlog   -loglang  eng  -serverip  nats://nats.newhorizons3000.org:4222  -logpattern  [ERR]  -logalias  LOGALIAS
#- serverip - NATS nats://xxxxx.yyy:port
# -logalias - make unique for each instance, become DEVICE.device in NATS
go build nhlog.go
cat nhlog.go | ./nhlog  -loglang eng -logurl http://127.0.0.1:8080/config -serverip nats://192.168.0.5:4222 -logpattern import -logalias loggertest
