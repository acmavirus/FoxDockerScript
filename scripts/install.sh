#!/bin/bash
# Copyright by AcmaTvirus

# Màu sắc cho output
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

echo -e "${GREEN}========== ĐANG CÀI ĐẶT FOXDOCKER PANEL ==========${NC}"

# 1. Kiểm tra quyền root
if [ "$EUID" -ne 0 ]; then
  echo -e "${RED}Vui lòng chạy script với quyền root (sudo).${NC}"
  exit 1
fi

# 2. Cài đặt Docker & Docker Compose nếu chưa có
if ! command -v docker &> /dev/null; then
    echo "Đang cài đặt Docker..."
    curl -fsSL https://get.docker.com | sh
    systemctl enable --now docker
fi

if ! command -v docker compose &> /dev/null; then
    echo "Đang cài đặt Docker Compose..."
    apt-get update && apt-get install -y docker-compose-plugin
fi

# 3. Tạo thư mục hệ thống
echo "Khởi tạo thư mục hệ thống tại /opt/foxdocker..."
mkdir -p /opt/foxdocker/data /opt/foxdocker/projects /opt/foxdocker/letsencrypt

# 4. Tải file cấu hình docker-compose.panel.yml từ nguồn (giả định GitHub)
# Ở đây chúng ta sẽ tạo file tại chỗ cho hướng dẫn này
cat <<EOF > /opt/foxdocker/docker-compose.yml
services:
  traefik:
    image: traefik:v3.0
    command:
      - "--providers.docker=true"
      - "--entrypoints.web.address=:80"
      - "--entrypoints.websecure.address=:443"
      - "--certificatesresolvers.myresolver.acme.email=admin@yourdomain.com"
      - "--certificatesresolvers.myresolver.acme.storage=/letsencrypt/acme.json"
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - "/var/run/docker.sock:/var/run/docker.sock:ro"
      - "/opt/foxdocker/letsencrypt:/letsencrypt"
    networks:
      - fox-net

  fox-admin:
    image: foxdocker/panel:latest
    container_name: fox-admin
    restart: always
    volumes:
      - "/var/run/docker.sock:/var/run/docker.sock"
      - "/opt/foxdocker/data:/app/data"
      - "/opt/foxdocker/projects:/projects"
    environment:
      - JWT_SECRET=$(openssl rand -hex 32)
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.fox-admin.rule=Host(\`panel.yourdomain.com\`)"
      - "traefik.http.routers.fox-admin.entrypoints=web"
    networks:
      - fox-net

networks:
  fox-net:
    driver: bridge
EOF

# 5. Khởi chạy Panel
echo "Đang khởi chạy FoxDocker Panel..."
cd /opt/foxdocker && docker compose up -d

echo -e "${GREEN}========== CÀI ĐẶT HOÀN TẤT ==========${NC}"
echo "Truy cập Panel tại cổng 8888 hoặc cấu hình Domain của bạn."
