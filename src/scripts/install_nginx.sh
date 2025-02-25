#!/bin/bash
# Cài đặt Nginx
sudo apt update
sudo apt install nginx -y
sudo systemctl start nginx
sudo systemctl enable nginx
echo "Nginx đã được cài đặt và khởi động."