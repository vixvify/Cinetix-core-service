## เริ่มใช้งานในเครื่อง

### 1. ติดตั้ง Dependency

```
go mod tidy
```

---

### 2. ตั้งค่า Environment

```
cp .env.example .env
```

แก้ไขค่าในไฟล์ `.env`

```
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=
DB_SSLMODE=disable
DB_TIMEZONE=Asia/Bangkok

APP_ENV=LOCAL

API_URL=/api/v1
PORT=8080

CORS_ALLOW_ORIGINS=http://localhost:3000

JWT_SECRET=
```

---

### 3. รัน Server

```
go run cmd/server/main.go
```

---

การรันด้วย Docker (Backend + Database)

1. ตั้งค่า Environment สำหรับ Docker

```
DB_HOST=db
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=login-database
DB_PORT=5432
DB_SSLMODE=disable
DB_TIMEZONE=Asia/Bangkok

APP_ENV=LOCAL

API_URL=/api/v1
PORT=8080

CORS_ALLOW_ORIGINS=http://localhost:3000

JWT_SECRET=
```

2. Run Container

```bash
docker compose up --build
```

3. Stop Container

```bash
docker compose down
```
