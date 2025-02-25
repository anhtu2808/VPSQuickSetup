#!/bin/bash
echo "Bắt đầu chạy read_ip.sh..."
hostname -I | awk '{print $1}'
echo "Kết thúc chạy read_ip.sh"
