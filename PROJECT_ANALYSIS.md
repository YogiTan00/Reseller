# Reseller - Sales Management System

## 📋 Overview

**Reseller** adalah sistem manajemen penjualan yang dibangun dengan Go menggunakan arsitektur Clean Architecture dan Microservices. Sistem ini menyediakan solusi lengkap untuk mengelola produk dan transaksi penjualan dengan dukungan HTTP REST API dan gRPC.

## 🏗️ Arsitektur

Project ini menggunakan **Clean Architecture** dengan pemisahan layer yang jelas:

```
┌─────────────────────────────────────────────────────────────┐
│                    Application Layer                         │
│                      (cmd/)                                  │
│                  Main Entry Points                           │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                    Delivery Layer                            │
│                  (HTTP/gRPC Handlers)                       │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                    Use Case Layer                            │
│                 (Business Logic)                             │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                    Repository Layer                          │
│                 (Data Access)                                │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                    Domain Layer                              │
│              (Entities & Interfaces)                         │
└─────────────────────────────────────────────────────────────┘
```

## 📁 Struktur Project

```
Reseller/
├── cmd/                        # Main application entry point
│   └── main.go                # Application bootstrap
├── config/                    # Configuration files
│   ├── config.go              # Application configuration
│   ├── database/              # Database setup
│   │   └── mysql/             # MySQL connection & migrations
│   └── html/                  # HTML configuration
├── pkg/                       # Shared packages
│   ├── exceptions/            # Custom error handling
│   ├── logger/                # Logging utilities (logrus)
│   └── utils/                 # Utility functions
├── services/                  # Microservices
│   ├── product/               # Product service
│   │   ├── domain/           # Entities & interfaces
│   │   ├── delivery/         # HTTP/gRPC handlers
│   │   ├── repository/       # Data access layer
│   │   └── usecase/          # Business logic
│   ├── transactions/          # Transaction service
│   │   ├── domain/           # Entities & interfaces
│   │   ├── delivery/         # HTTP/gRPC handlers
│   │   ├── repository/       # Data access layer
│   │   └── usecase/          # Business logic
│   └── html/                 # Web interface assets
│       └── form/             # HTML forms for CRUD operations
├── proto/                     # Protocol buffer definitions
│   └── buffer/               # Generated protobuf files
├── .env.example              # Environment variables template
├── .gitignore               # Git ignore rules
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
└── buf.gen.yaml             # Buf configuration for protobuf
```

## 🛠️ Tech Stack

### Core Technologies
| Technology | Versi | Penggunaan |
|------------|-------|------------|
| Go | 1.23.0 | Bahasa pemrograman utama |
| GORM | - | ORM untuk operasi database |
| MySQL | - | Database utama |
| gRPC | - | RPC framework untuk komunikasi internal |
| gRPC Gateway | - | HTTP ke gRPC proxy |
| Gorilla Mux | - | HTTP router |

### Supporting Libraries
| Library | Fungsi |
|---------|--------|
| logrus | Structured logging |
| godotenv | Environment variable management |
| golang-migrate | Database migrations |
| mockery | Mock generation untuk testing |

### Development Tools
- **protoc** - Protocol buffer compiler
- **buf** - Protocol buffer management
- **Git** - Version control

## 🚀 Fitur Utama

### 1. Product Service
Manajemen katalog produk dengan operasi CRUD lengkap:

#### Endpoint HTTP:
- `POST /api/v1/product/create` - Buat produk baru
- `POST /api/v1/product/update` - Update produk existing
- `POST /api/v1/product/delete/{id}` - Hapus produk (soft delete)
- `GET /api/v1/product/view/{id}` - Lihat detail produk
- `GET /api/v1/product/list` - List produk dengan filter

#### Model Produk:
```go
type Product struct {
    ID           string  `gorm:"primaryKey"`
    Name         string
    TypeSize     string
    Image        string
    DefaultPrice float64
    Price        float64
    CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    *time.Time  // Soft delete
}
```

### 2. Transaction Service
Manajemen transaksi penjualan:

#### Endpoint HTTP:
- `POST /api/v1/transaction/create` - Buat transaksi tunggal
- `POST /api/v1/transaction/create/multiple` - Buat multiple transaksi
- `POST /api/v1/transaction/update` - Update transaksi
- `POST /api/v1/transaction/delete/{id}` - Hapus transaksi
- `GET /api/v1/transaction/view/{id}` - Lihat detail transaksi
- `GET /api/v1/transaction/list` - List transaksi dengan filter tanggal

#### Model Transaksi:
```go
type Transaction struct {
    ID           string    `gorm:"primaryKey"`
    IDName       string
    ReturnTrans  bool
    SalesDate    time.Time
    Unit         int
    Description  string
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
```

### 3. HTML Web Interface
Antarmuka web user-friendly dengan Bootstrap:
- Form untuk create product & transaction
- List view untuk products & transactions
- Support untuk multiple transaction creation
- Responsive design

## 💾 Database Schema

### Products Table
| Column | Type | Description |
|--------|------|-------------|
| id | VARCHAR | Primary key |
| name | VARCHAR | Nama produk |
| type_size | VARCHAR | Tipe/ukuran produk |
| image | VARCHAR | URL gambar produk |
| default_price | DECIMAL | Harga default |
| price | DECIMAL | Harga saat ini |
| created_at | TIMESTAMP | Waktu pembuatan |
| updated_at | TIMESTAMP | Waktu update |
| deleted_at | TIMESTAMP | Soft delete (nullable) |

### Transactions Table
| Column | Type | Description |
|--------|------|-------------|
| id | VARCHAR | Primary key |
| id_name | VARCHAR | Nama/ID referensi |
| return_trans | BOOLEAN | Status return |
| sales_date | DATE | Tanggal penjualan |
| unit | INT | Jumlah unit |
| description | TEXT | Deskripsi |
| created_at | TIMESTAMP | Waktu pembuatan |
| updated_at | TIMESTAMP | Waktu update |

## 🔧 Cara Menjalankan

### Prerequisites:
1. Go 1.23.0 atau higher
2. MySQL Server
3. Protocol Buffer compiler (protoc)
4. Buf

### Instalasi:
```bash
# Clone repository
git clone <repository-url>
cd Reseller

# Install dependencies
go mod download

# Setup environment variables
cp .env.example .env
# Edit .env sesuai konfigurasi

# Run database migrations
# (Setup database MySQL sesuai config)

# Generate protobuf files
buf generate

# Run application
go run cmd/main.go
```

### Environment Variables:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=reseller_db
SERVER_PORT=8080
GRPC_PORT=50051
```

## 📊 Status Implementasi

### ✅ Fitur yang Sudah Selesai:
- [x] CRUD lengkap untuk Product & Transaction
- [x] gRPC Services dengan HTTP Gateway
- [x] Web Interface dengan Bootstrap
- [x] Database Migrations
- [x] Date filtering untuk transactions
- [x] Multiple transaction creation
- [x] Soft delete support
- [x] Logging system dengan request tracking
- [x] Clean Architecture implementation

### 🔄 Update Terbaru:
- Upgrade ke Go 1.23.0
- Update Buf version
- Enhanced date filtering pada transaction list
- Added HTML forms untuk transaction creation
- Fixed pagination dan date range filtering

## 🔍 Poin-poin Perbaikan yang Direkomendasikan

### 🔴 Prioritas Tinggi (Critical)

#### 1. **Testing Coverage**
**Masalah:** Tidak ada file test sama sekali
**Dampak:** High risk untuk bugs, sulit untuk refactor
**Rekomendasi:**
- Tambah unit tests untuk use case layer
- Tambah integration tests untuk API endpoints
- Gunakan testify atau testing package dari Go
- Target coverage minimal 70%

**Contoh implementasi:**
```go
// services/product/usecase/product_usecase_test.go
func TestProductUseCase_Create(t *testing.T) {
    // Mock repository
    mockRepo := new(mocks.ProductRepository)
    useCase := NewProductUseCase(mockRepo)

    // Test case
    mockRepo.On("Create", mock.Anything).Return(nil)
    err := useCase.Create(&domain.Product{})

    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
}
```

#### 2. **API Documentation**
**Masalah:** Tidak ada dokumentasi API (Swagger/OpenAPI)
**Dampak:** Sulit untuk developer baru, integrasi API manual
**Rekomendasi:**
- Integrasikan Swaggo untuk Swagger documentation
- Tambah annotations pada semua handlers
- Generate dan host Swagger UI
- Document semua request/response schemas

**Implementasi:**
```go
// @title Reseller API
// @version 1.0
// @description Sales Management System API
// @host localhost:8080
// @BasePath /api/v1

// @Summary Create Product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body domain.Product true "Product object"
// @Success 200 {object} domain.Product
// @Router /product/create [post]
```

#### 3. **Authentication & Authorization**
**Masalah:** Tidak ada layer keamanan sama sekali
**Dampak:** API terbuka untuk semua orang, riskan
**Rekomendasi:**
- Implement JWT authentication
- Tambah middleware untuk auth check
- Implement role-based access control (RBAC)
- Tambah rate limiting untuk prevent DDoS
- Implement CORS configuration

**Contoh middleware:**
```go
func AuthMiddleware(jwtSecret string) mux.MiddlewareFunc {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            token := r.Header.Get("Authorization")
            // Validate JWT token
            // Set user context
            next.ServeHTTP(w, r)
        })
    }
}
```

### 🟡 Prioritas Sedang (Important)

#### 4. **Input Validation**
**Masalah:** Validasi input masih basic
**Dampak:** Invalid data bisa masuk ke database
**Rekomendasi:**
- Gunakan validator library (go-playground/validator)
- Implement custom validation rules
- Return structured error messages
- Sanitize user input untuk prevent injection

**Implementasi:**
```go
type Product struct {
    Name     string `json:"name" validate:"required,min=3,max=100"`
    Price    float64 `json:"price" validate:"required,gt=0"`
    TypeSize string `json:"type_size" validate:"required,alpha"`
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
    var product domain.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        // Handle error
    }

    if err := h.validator.Struct(&product); err != nil {
        // Return validation errors
    }
}
```

#### 5. **Error Handling**
**Masalah:** Error handling masih basic, tidak ada standard error response
**Dampak:** Sulit untuk debug di production, user experience kurang baik
**Rekomendasi:**
- Buat standard error response format
- Implement error codes dan categories
- Log error dengan context yang lengkap
- Return user-friendly error messages
- Handle panic dengan recovery middleware

**Format error response:**
```go
type ErrorResponse struct {
    Error struct {
        Code    string `json:"code"`
        Message string `json:"message"`
        Details string `json:"details,omitempty"`
    } `json:"error"`
    Timestamp string `json:"timestamp"`
    Path      string `json:"path"`
}

// Contoh usage:
return ErrorResponse{
    Error: struct{
        Code: "VALIDATION_ERROR",
        Message: "Invalid input",
        Details: "Name is required",
    },
    Timestamp: time.Now().Format(time.RFC3339),
    Path: r.URL.Path,
}
```

#### 6. **Configuration Management**
**Masalah:** Configuration scattered, tidak ada centralization
**Dampak:** Sulit untuk manage environment yang berbeda
**Rekomendasi:**
- Gunakan Viper untuk configuration management
- Support multiple config sources (file, env, flags)
- Implement config validation
- Separate config per environment (dev, staging, prod)

#### 7. **Caching Layer**
**Masalah:** Tidak ada caching, semua request hit database
**Dampak:** Performance issue untuk high traffic
**Rekomendasi:**
- Integrasikan Redis untuk caching
- Cache frequently accessed data (product lists)
- Implement cache invalidation strategy
- Tambah cache headers untuk HTTP responses

**Implementasi:**
```go
type ProductUseCase struct {
    repo    domain.ProductRepository
    cache   CacheInterface
}

func (uc *ProductUseCase) GetByID(id string) (*domain.Product, error) {
    // Check cache first
    cached, err := uc.cache.Get(fmt.Sprintf("product:%s", id))
    if err == nil {
        return cached, nil
    }

    // If not in cache, get from DB
    product, err := uc.repo.GetByID(id)
    if err != nil {
        return nil, err
    }

    // Set cache
    uc.cache.Set(fmt.Sprintf("product:%s", id), product, 5*time.Minute)
    return product, nil
}
```

### 🟢 Prioritas Rendah (Nice to Have)

#### 8. **Monitoring & Health Checks**
**Masalah:** Tidak ada monitoring, health checks, atau metrics
**Dampak:** Sulit untuk monitoring di production
**Rekomendasi:**
- Tambah health check endpoint
- Implement Prometheus metrics
- Monitor database connection pool
- Track request metrics (latency, throughput)
- Setup alerting untuk critical errors

**Health check endpoint:**
```go
func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
    status := struct {
        Status   string `json:"status"`
        Database string `json:"database"`
        Cache    string `json:"cache"`
    }{
        Status: "healthy",
    }

    // Check DB connection
    if err := h.db.Ping(); err != nil {
        status.Database = "unhealthy"
        status.Status = "degraded"
    } else {
        status.Database = "healthy"
    }

    json.NewEncoder(w).Encode(status)
}
```

#### 9. **API Versioning Strategy**
**Masalah:** Tidak ada strategy untuk API versioning
**Dampak:** Breaking changes akan affect semua clients
**Rekomendasi:**
- Implement URL-based versioning (/v1, /v2)
- Document deprecation timeline
- Support multiple versions simultaneously
- Communication plan untuk API changes

#### 10. **Database Connection Pooling**
**Masalah:** Tidak ada explicit connection pool configuration
**Dampak:** Potensi connection exhaustion di high load
**Rekomendasi:**
- Configure connection pool parameters
- Monitor connection pool metrics
- Implement connection timeout
- Handle connection errors gracefully

**Configuration:**
```go
db.DB().SetMaxOpenConns(100)
db.DB().SetMaxIdleConns(10)
db.DB().SetConnMaxLifetime(time.Hour)
db.DB().SetConnMaxIdleTime(5 * time.Minute)
```

#### 11. **Logging Enhancement**
**Masalah:** Logging masih basic, kurang context
**Dampak:** Sulit untuk debugging di production
**Rekomendasi:**
- Tambah request ID untuk tracing
- Log dengan structured fields
- Implement log levels (debug, info, warn, error)
- Log sensitive data dengan masking
- Integrate dengan ELK/Loki untuk log aggregation

**Enhanced logging:**
```go
logger.WithFields(logrus.Fields{
    "request_id": requestID,
    "method": r.Method,
    "path": r.URL.Path,
    "duration": duration.Milliseconds(),
    "status": statusCode,
    "user_id": userID,
}).Info("Request completed")
```

#### 12. **Database Indexing**
**Masalah:** Tidak jelas apakah ada indexing strategy
**Dampak:** Query performance akan degrade seiring bertambahnya data
**Rekomendasi:**
- Add indexes pada frequently queried columns
- Create composite indexes untuk complex queries
- Monitor query performance dengan EXPLAIN
- Regular index maintenance

**Migration untuk indexing:**
```sql
CREATE INDEX idx_products_deleted_at ON products(deleted_at);
CREATE INDEX idx_transactions_sales_date ON transactions(sales_date);
CREATE INDEX idx_transactions_id_name ON transactions(id_name);
```

#### 13. **Pagination Implementation**
**Masalah:** List endpoints tidak ada pagination
**Dampak:** Performance issue untuk large datasets
**Rekomendasi:**
- Implement pagination untuk semua list endpoints
- Support page size limiting
- Return pagination metadata (total, pages, current)
- Implement cursor-based pagination untuk better performance

**Request/response format:**
```go
type ListRequest struct {
    Page     int    `json:"page" validate:"min=1"`
    PageSize int    `json:"page_size" validate:"min=1,max=100"`
}

type ListResponse struct {
    Data       []Product `json:"data"`
    Pagination struct {
        CurrentPage int `json:"current_page"`
        PageSize    int `json:"page_size"`
        TotalPages  int `json:"total_pages"`
        TotalItems  int `json:"total_items"`
    } `json:"pagination"`
}
```

#### 14. **Transaction Management**
**Masalah:** Tidak ada explicit transaction management untuk multi-step operations
**Dampak:** Data inconsistency jika salah satu step gagal
**Rekomendasi:**
- Implement database transactions untuk critical operations
- Handle transaction rollback on errors
- Implement distributed transactions untuk cross-service operations

**Contoh:**
```go
func (uc *TransactionUseCase) CreateMultiple(transactions []domain.Transaction) error {
    tx := uc.db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    for _, trans := range transactions {
        if err := tx.Create(&trans).Error; err != nil {
            tx.Rollback()
            return err
        }
    }

    return tx.Commit().Error
}
```

#### 15. **Dockerization**
**Masalah:** Tidak ada Docker setup
**Dampak:** Sulit untuk deployment, environment inconsistency
**Rekomendasi:**
- Create Dockerfile untuk application
- Create docker-compose untuk local development
- Multi-stage build untuk smaller image size
- Document Docker commands di README

**Dockerfile example:**
```dockerfile
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

## 📝 Kesimpulan

Project **Reseller** adalah sistem manajemen penjualan yang well-structured dengan arsitektur Clean Architecture dan pemisahan microservices yang baik. Namun, masih banyak area yang perlu diperbaiki untuk production readiness:

### Quick Wins (bisa dikerjakan 1-2 minggu):
1. Tambah unit tests untuk critical paths
2. Implement basic authentication
3. Add input validation
4. Setup Swagger documentation

### Medium Term (1-2 bulan):
1. Complete testing coverage
2. Implement caching
3. Enhanced error handling
4. Add monitoring & health checks

### Long Term (3+ bulan):
1. Complete CI/CD pipeline
2. Advanced security features
3. Performance optimization
4. Scalability improvements

### Priority Order Rekomendasi:
1. **Testing** - Fundamental untuk code quality
2. **Authentication** - Critical untuk security
3. **Documentation** - Essential untuk maintainability
4. **Error Handling** - Important untuk UX
5. **Validation** - Important untuk data integrity
6. **Caching** - Performance improvement
7. **Monitoring** - Operational visibility
8. **Docker** - Deployment consistency

---

**Dokumentasi ini dibuat pada:** 2026-02-18
**Versi Dokumen:** 1.0
**Last Updated:** 2026-02-18
