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

# Tìm hoặc cài đặt Docker Compose
COMPOSE_CMD=$(find_compose_cmd)

if [ -z "$COMPOSE_CMD" ]; then
    echo "Đang cài đặt Docker Compose..."
    # Thử cài qua apt trước (cho Debian/Ubuntu)
    if command -v apt-get &> /dev/null; then
        apt-get update && apt-get install -y docker-compose-plugin
    fi
    
    COMPOSE_CMD=$(find_compose_cmd)
    
    # Nếu vẫn chưa có, tải binary trực tiếp (Fallback cho CentOS/Legacy)
    if [ -z "$COMPOSE_CMD" ]; then
        echo "Cài đặt plugin thất bại, chuyển sang tải binary..."
        curl -SL https://github.com/docker/compose/releases/download/v2.24.6/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose
        chmod +x /usr/local/bin/docker-compose
        COMPOSE_CMD="docker-compose"
    fi
fi

echo -e "Sử dụng lệnh Compose: ${GREEN}$COMPOSE_CMD${NC}"

# ... (Giữ nguyên logic kiểm tra cổng)

# 3. Kiểm tra cổng khả dụng (Mặc định 8888)
PANEL_PORT=8888
check_port() {
    # Kiểm tra bằng ss hoặc netstat tùy hệ điều hành
    if command -v ss &> /dev/null; then
        ss -tuln | grep -q ":$1 "
    else
        netstat -tuln | grep -q ":$1 "
    fi
}

if check_port $PANEL_PORT; then
    echo -e "${RED}Cổng $PANEL_PORT đã bị chiếm dụng. Đang tìm một cổng trống ngẫu nhiên...${NC}"
    while :
    do
        # Sinh cổng ngẫu nhiên trong khoảng 10000 - 60000
        PANEL_PORT=$(( ( RANDOM % 50000 )  + 10000 ))
        if ! check_port $PANEL_PORT; then
            break
        fi
    done
    echo -e "${GREEN}FoxDocker sẽ sử dụng cổng mới: $PANEL_PORT${NC}"
fi

# 4. Tạo thư mục hệ thống
echo "Khởi tạo thư mục hệ thống tại /opt/foxdocker..."
mkdir -p /opt/foxdocker/data /opt/foxdocker/projects /opt/foxdocker/letsencrypt /opt/foxdocker/scripts

# 4.1 Tải script cập nhật
echo "Đang tải script cập nhật..."
curl -sSL https://raw.githubusercontent.com/acmavirus/FoxDockerScript/main/scripts/update.sh -o /opt/foxdocker/scripts/update.sh
chmod +x /opt/foxdocker/scripts/update.sh

# 5. Tạo file cấu hình docker-compose.yml
echo "Đang tạo cấu hình Docker Compose..."
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
    image: ghcr.io/acmavirus/foxdocker-panel:latest
    container_name: fox-admin
    restart: always
    ports:
      - "$PANEL_PORT:8888"
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

# 6. Khởi chạy Panel
echo "Đang tải và khởi chạy FoxDocker Panel từ GitHub Packages..."
cd /opt/foxdocker && $COMPOSE_CMD pull && $COMPOSE_CMD up -d

echo -e "${GREEN}========== CÀI ĐẶT HOÀN TẤT ==========${NC}"
echo -e "Truy cập Panel tại: ${GREEN}http://IP_CUA_BAN:$PANEL_PORT${NC}"
echo -e "Hoặc cấu hình Domain trong file /opt/foxdocker/docker-compose.yml"
