#openssl genrsa -out server.key 2048
#openssl ecparam -genkey -name secp384r1 -out server.key
#openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout server.pem -out cert.pem -config ssl.conf
