cp /home/oem/go/src/github.com/nh3000-org/radio/db/radio.sql /home/oem/go/src/github.com/nh3000-org/radio/cmd/db/radio.sql.backup
sudo -u postgres  pg_dump radio > /home/oem/go/src/github.com/nh3000-org/radio/cmd/db/radio.sql.before
sudo -u postgres psql -c "drop database radio"
sudo -u postgres psql -c "create database radio"
sudo -u postgres psql -h localhost -p 5432 -U postgres radio < dbcreate.sql
sudo -u postgres psql -c "\l"
sudo -u postgres  pg_dump radio > /home/oem/go/src/github.com/nh3000-org/radio/cmd/db/radio.sql
