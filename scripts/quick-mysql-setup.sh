#!/bin/bash

# Quick MySQL Setup Script

echo "🐬 Quick MySQL Setup untuk Perpustakaan API"
echo ""

# Default values
DB_NAME="belajargo"
DB_ROOT_PASSWORD="1"

echo "Membuat database $DB_NAME..."

mysql -u root -p"$DB_ROOT_PASSWORD" -e "
CREATE DATABASE IF NOT EXISTS $DB_NAME CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
SHOW DATABASES;
USE $DB_NAME;
SELECT 'Database $DB_NAME siap digunakan!' as Status;
"

if [ $? -eq 0 ]; then
    echo "✅ Database berhasil dibuat!"
    echo ""
    echo "🚀 Langkah selanjutnya:"
    echo "1. Pastikan .env sudah dikonfigurasi dengan benar"
    echo "2. Jalankan: make run"
    echo "3. Auto migration akan berjalan otomatis"
else
    echo "❌ Gagal membuat database!"
    echo "Pastikan MySQL sudah running dan password root benar"
fi
