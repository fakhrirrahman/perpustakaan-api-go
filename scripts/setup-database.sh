#!/bin/bash

# Database Setup Script untuk Perpustakaan API - MySQL

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== Database Setup untuk Perpustakaan API (MySQL) ===${NC}"
echo ""

# Check if MySQL is installed
if ! command -v mysql &> /dev/null; then
    echo -e "${RED}MySQL tidak ditemukan!${NC}"
    echo -e "${YELLOW}Install MySQL terlebih dahulu:${NC}"
    echo "Ubuntu/Debian: sudo apt-get install mysql-server"
    echo "macOS: brew install mysql"
    echo "Windows: Download dari https://dev.mysql.com/downloads/mysql/"
    exit 1
fi

echo -e "${GREEN}MySQL ditemukan!${NC}"

# Database configuration
DB_NAME="belajargo"
DB_USER="root"
DB_PASSWORD="1"
DB_HOST="localhost"
DB_PORT="3306"

echo -e "${BLUE}Konfigurasi Database:${NC}"
echo "Database: $DB_NAME"
echo "User: $DB_USER"
echo "Host: $DB_HOST"
echo "Port: $DB_PORT"
echo ""

# Ask for confirmation
read -p "Lanjutkan dengan konfigurasi ini? (y/N): " confirm
if [ "$confirm" != "y" ] && [ "$confirm" != "Y" ]; then
    echo -e "${YELLOW}Setup dibatalkan.${NC}"
    exit 0
fi

echo -e "${BLUE}Membuat database dan user...${NC}"

# Create database and user
mysql -u root -p"$DB_PASSWORD" << EOF
-- Create database
CREATE DATABASE IF NOT EXISTS $DB_NAME CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Create user (optional if not using root)
-- CREATE USER IF NOT EXISTS '$DB_USER'@'localhost' IDENTIFIED BY '$DB_PASSWORD';

-- Grant privileges
-- GRANT ALL PRIVILEGES ON $DB_NAME.* TO '$DB_USER'@'localhost';

-- Use database
USE $DB_NAME;

-- Show database
SHOW DATABASES;

QUIT;
EOF

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Database setup berhasil!${NC}"
else
    echo -e "${RED}✗ Database setup gagal!${NC}"
    exit 1
fi

# Create .env file if not exists
ENV_FILE=".env"
if [ ! -f "$ENV_FILE" ]; then
    echo -e "${BLUE}Membuat file .env...${NC}"
    cat > $ENV_FILE << EOF
# Database Configuration
DB_HOST=$DB_HOST
DB_PORT=$DB_PORT
DB_USER=$DB_USER
DB_PASS=$DB_PASSWORD
DB_NAME=$DB_NAME
DB_TZ=Asia/Jakarta

# Server Configuration
SERVER_HOST=0.0.0.0
SERVER_PORT=8080

# JWT Configuration
JWT_KEY=your-super-secret-jwt-key-change-this-in-production

EOF
    echo -e "${GREEN}✓ File .env berhasil dibuat!${NC}"
else
    echo -e "${YELLOW}File .env sudah ada, skip pembuatan.${NC}"
fi

echo ""
echo -e "${GREEN}=== Setup Selesai! ===${NC}"
echo ""
echo -e "${BLUE}Langkah selanjutnya:${NC}"
echo "1. Jalankan aplikasi: ${YELLOW}make run${NC}"
echo "2. Atau untuk development: ${YELLOW}make run-dev${NC}"
echo "3. Database akan otomatis di-migrate saat aplikasi start"
echo ""
echo -e "${BLUE}Test database connection:${NC}"
echo "mysql -h $DB_HOST -P $DB_PORT -u $DB_USER -p $DB_NAME"
echo ""
echo -e "${BLUE}API Endpoints yang tersedia:${NC}"
echo "GET    /api/authors      - List semua penulis"
echo "GET    /api/authors/:id  - Detail penulis"
echo "POST   /api/authors      - Buat penulis baru (perlu JWT)"
echo "PUT    /api/authors/:id  - Update penulis (perlu JWT)"
echo "DELETE /api/authors/:id  - Hapus penulis (perlu JWT)"
