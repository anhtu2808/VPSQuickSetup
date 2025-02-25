#!/bin/bash
# Tạo file docker-compose
cat <<EOL > docker-compose.yml
version: '3'
services:
  app:
    image: nginx
    ports:
      - "80:80"
EOL
echo "File docker-compose.yml đã được tạo."