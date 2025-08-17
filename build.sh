#!/usr/bin/bash

docker build -t todos-app ./service 
rm ~/Documents/images/server-image.tar || true
mkdir -pv ~/Documents/images/
docker save -o ~/Documents/images/server-image.tar todos-app:latest
ls -lh ~/Documents/images/server-image.tar
tar tf ~/Documents/images/server-image.tar
