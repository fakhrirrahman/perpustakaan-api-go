# 📚 API Testing Commands

## Base URL
```
BASE_URL=http://localhost:8080/api
```

## Authors API

### 1. Get All Authors
```bash
curl -X GET "$BASE_URL/authors" \
  -H "Content-Type: application/json"
```

### 2. Get Author by ID
```bash
curl -X GET "$BASE_URL/authors/123e4567-e89b-12d3-a456-426614174000" \
  -H "Content-Type: application/json"
```

### 3. Create New Author (Need JWT Token)
```bash
curl -X POST "$BASE_URL/authors" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "J.K. Rowling",
    "email": "jk.rowling@example.com",
    "bio": "British author, best known for the Harry Potter series"
  }'
```

### 4. Update Author (Need JWT Token)
```bash
curl -X PUT "$BASE_URL/authors/123e4567-e89b-12d3-a456-426614174000" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "J.K. Rowling Updated",
    "email": "jk.rowling.updated@example.com",
    "bio": "Updated biography"
  }'
```

### 5. Delete Author (Need JWT Token)
```bash
curl -X DELETE "$BASE_URL/authors/123e4567-e89b-12d3-a456-426614174000" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Testing Validation Errors

### Invalid Email
```bash
curl -X POST "$BASE_URL/authors" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Test Author",
    "email": "invalid-email",
    "bio": "Test bio"
  }'
```

### Missing Required Fields
```bash
curl -X POST "$BASE_URL/authors" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "bio": "Test bio only"
  }'
```

### Too Short Name
```bash
curl -X POST "$BASE_URL/authors" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "A",
    "email": "short@example.com",
    "bio": "Test bio"
  }'
```

## Expected Responses

### Success Response
```json
{
  "code": 200,
  "message": "Success message",
  "data": {...}
}
```

### Error Response
```json
{
  "code": 99,
  "message": "Error message",
  "data": ""
}
```

## Authentication

Untuk endpoints yang memerlukan authentication, gunakan Bearer token:
```
Authorization: Bearer <your-jwt-token>
```

## Swagger Documentation

Akses dokumentasi interaktif di:
```
http://localhost:8080/swagger/index.html
```
