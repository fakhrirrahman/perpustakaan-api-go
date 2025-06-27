# 📚 Perpustakaan API - Go

API untuk sistem perpustakaan menggunakan Go, Fiber, GORM, dan PostgreSQL dengan Clean Architecture.

## 🚀 Quick Start

### Prerequisites
- Go 1.24.1+
- MySQL 8.0+
- Make (optional, untuk menjalankan commands dengan mudah)

### 1. Clone & Setup
```bash
git clone <repository-url>
cd perpustakaan-api-go
make setup  # Install dependencies dan tools
```

### 2. Database Setup
```bash
# Setup database MySQL otomatis
./scripts/setup-database.sh

# Atau manual setup database, kemudian copy .env
cp .env.example .env
# Edit .env sesuai konfigurasi database Anda
```

### 3. Jalankan Aplikasi
```bash
# Development dengan auto-reload
make run-dev

# Atau jalankan biasa
make run

# Atau tanpa Make
go run main.go
```

## 📁 Struktur Arsitektur

Proyek ini menggunakan **Clean Architecture** dengan layer-layer berikut:

```
📁 perpustakaan-api-go/
├── 📁 domain/          # Business entities & interfaces
│   ├── author.go       # Author entity & contracts
│   ├── book.go         # Book entity & contracts
│   ├── customer.go     # Customer entity & contracts
│   └── error.go        # Custom error definitions
├── 📁 dto/             # Data Transfer Objects
│   ├── author_data.go  # Author request/response
│   ├── book_data.go    # Book request/response
│   └── respone.go      # Generic response wrapper
├── 📁 internal/
│   ├── 📁 api/         # HTTP handlers (Controllers)
│   │   ├── author.go   # Author HTTP handlers
│   │   └── book.go     # Book HTTP handlers
│   ├── 📁 service/     # Business logic layer
│   │   ├── author.go   # Author business logic
│   │   └── book.go     # Book business logic  
│   ├── 📁 repository/  # Data access layer
│   │   ├── author.go   # Author database operations
│   │   └── book.go     # Book database operations
│   ├── 📁 config/      # Configuration management
│   ├── 📁 connection/  # Database connections
│   └── 📁 util/        # Utility functions
├── 📁 scripts/         # Automation scripts
└── main.go             # Application entry point
```

### Alur Data Flow:
```
HTTP Request → API Layer → Service Layer → Repository Layer → Database
HTTP Response ← API Layer ← Service Layer ← Repository Layer ← Database
```

## 🔧 Available Commands

Gunakan `make help` untuk melihat semua commands yang tersedia:

```bash
make install     # Install dependencies
make run         # Jalankan aplikasi
make run-dev     # Development dengan auto-reload
make build       # Build aplikasi
make test        # Jalankan tests
make format      # Format code
make lint        # Jalankan linter
make clean       # Bersihkan build files
```

## 📖 API Documentation

### Swagger Documentation
Akses dokumentasi API interaktif di: **http://localhost:8080/swagger/index.html**

### Generate Swagger Docs
```bash
# Generate dokumentasi
make swagger

# Generate docs dan jalankan aplikasi
make swagger-serve
```

### Authors Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/api/authors` | List semua penulis | ❌ |
| GET | `/api/authors/:id` | Detail penulis | ❌ |
| POST | `/api/authors` | Buat penulis baru | ✅ |
| PUT | `/api/authors/:id` | Update penulis | ✅ |
| DELETE | `/api/authors/:id` | Hapus penulis | ✅ |

### Request Examples

#### Create Author
```bash
curl -X POST http://localhost:8080/api/authors \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "J.K. Rowling",
    "email": "jk.rowling@example.com",
    "bio": "British author, best known for the Harry Potter series"
  }'
```

#### Get All Authors
```bash
curl http://localhost:8080/api/authors
```

#### Get Author by ID
```bash
curl http://localhost:8080/api/authors/uuid-here
```

#### Update Author
```bash
curl -X PUT http://localhost:8080/api/authors/uuid-here \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "J.K. Rowling Updated",
    "email": "jk.rowling.updated@example.com",
    "bio": "Updated bio"
  }'
```

## 🗄️ Database

### Author Table Schema
```sql
CREATE TABLE authors (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    bio TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### Migration
Database migration berjalan otomatis menggunakan GORM AutoMigrate saat aplikasi start.

## 🔐 Authentication

API menggunakan JWT untuk authentication. Beberapa endpoints memerlukan valid JWT token:

```bash
# Header format
Authorization: Bearer <your-jwt-token>
```

## 🐛 Error Handling

API mengembalikan error dalam format standar:

```json
{
  "code": 99,
  "message": "Error message here",
  "data": ""
}
```

## 🧪 Testing

```bash
# Jalankan semua tests
make test

# Test specific package
go test ./internal/service/...
```

## 📦 Dependencies

- **[Fiber](https://github.com/gofiber/fiber)** - HTTP framework
- **[GORM](https://gorm.io/)** - ORM library
- **[MySQL Driver](https://github.com/go-sql-driver/mysql)** - Database driver
- **[Validator](https://github.com/go-playground/validator)** - Input validation
- **[UUID](https://github.com/google/uuid)** - UUID generation
- **[JWT](https://github.com/golang-jwt/jwt)** - JWT authentication

## 🚀 Deployment

### Build Production
```bash
make build
./bin/perpustakaan-api
```

### Docker
```bash
make docker-build
make docker-run
```

## 🤝 Contributing

1. Fork repository
2. Buat feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open Pull Request

## 📝 License

This project is licensed under the MIT License.

## 📞 Support

Jika ada pertanyaan atau issues, silakan buat issue di repository ini.
