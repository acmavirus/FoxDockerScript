#!/bin/bash
# Copyright by AcmaTvirus

echo "Đang kiểm tra cập nhật cho FoxDocker Panel..."

cd /opt/foxdocker
docker compose pull
docker compose up -d

echo "Đã cập nhật FoxDocker Panel lên bản mới nhất."
