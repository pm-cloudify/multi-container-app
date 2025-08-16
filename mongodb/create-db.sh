#!/usr/bin/bash

if [[ $EUID -eq 0  ]]; then 
    echo "This script needs root access"
    exit 1
fi

mkdir -p  /opt/apps/mongodb/configdb  /opt/apps/mongodb/db
docker compose up -d
