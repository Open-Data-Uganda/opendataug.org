#!/bin/bash
sudo chmod 666 /var/run/docker.sock 

sudo sh ./down.sh

echo "Build starting......\n"

echo "Pulling ropes......\n"

git pull 

echo "Build starting......\n"

docker-compose up -d --build

echo "......\n"

echo "......\n"

docker images

docker system prune -a -f 

echo "......\n"

docker ps 