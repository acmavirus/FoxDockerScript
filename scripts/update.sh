#!/bin/bash
# Copyright by AcmaTvirus

echo "Đang kiểm tra cập nhật cho FoxDocker Panel..."

# Hàm tìm lệnh Docker Compose hợp lệ
find_compose_cmd() {
    if docker compose version &> /dev/null; then
        echo "docker compose"
    elif command -v docker-compose &> /dev/null; then
        echo "docker-compose"
    else
        return 1
    fi
}

COMPOSE_CMD=$(find_compose_cmd)

if [ -z "$COMPOSE_CMD" ]; then
    echo "Lỗi: Không tìm thấy Docker Compose. Vui lòng cài đặt Docker Compose trước."
    exit 1
fi

cd /opt/foxdocker
$COMPOSE_CMD pull
$COMPOSE_CMD up -d

echo "Đã cập nhật FoxDocker Panel lên bản mới nhất."
