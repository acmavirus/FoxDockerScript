# 🦊 FoxDocker Panel

**FoxDocker Panel** là một giải pháp quản trị Hosting hiện đại, siêu nhẹ và tối ưu cho Docker. Được xây dựng bằng **Go** và **Vue 3**, FoxDocker mang lại tốc độ vượt trội và trải nghiệm người dùng mượt mà theo phong cách tối giản.

## 🚀 Cài đặt nhanh (1-Click Install)

Chạy lệnh duy nhất sau trên VPS của bạn (hỗ trợ Ubuntu, Debian, CentOS):

```bash
curl -sSL https://raw.githubusercontent.com/acmavirus/FoxDockerScript/main/scripts/install.sh | bash
```

## ✨ Tính năng nổi bật

- ⚡ **Siêu nhẹ**: Khởi động trong 120 giây, chạy mượt trên VPS 512MB RAM.
- 🛡️ **Bảo mật**: Tích hợp quét lỗ hổng Image và thông báo qua Telegram/Discord.
- 🌐 **Traefik Zero-Config**: Tự động cấp phát SSL (Let's Encrypt) và quản lý Domain thông minh.
- 📊 **Live Health**: Theo dõi tài nguyên CPU/RAM/Disk theo thời gian thực.
- 📦 **App Store**: Cài đặt nhanh WordPress, Nginx, Redis... chỉ với 1-click.

## 🏗️ Cấu trúc dự án

- `cmd/fox-admin/`: Code nguồn chính của Backend (Go).
- `web/`: Giao diện người dùng (Vue 3, Vite, Tailwind CSS).
- `scripts/`: Chứa các kịch bản cài đặt và cập nhật.
- `templates/`: Các mẫu Docker Compose cho các ứng dụng phổ biến.

## 🛠️ Quản lý hệ thống

Sau khi cài đặt, bạn có thể quản lý Panel thông qua các lệnh Docker tiêu chuẩn:

```bash
# Xem logs
docker logs -f fox-admin

# Cập nhật bản mới
bash /opt/foxdocker/scripts/update.sh

# Restart hệ thống
cd /opt/foxdocker && docker compose restart
```

---

> **Copyright by AcmaTvirus** 🦊
