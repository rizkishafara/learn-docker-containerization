# 🧩 UTS_PSS – Web App (Golang + Fiber + MySQL + Docker)

Proyek ini merupakan aplikasi web berbasis **Golang (Fiber Framework)** dengan integrasi **MySQL Database**, dijalankan menggunakan **Docker Compose**.  

---

## ⚙️ Arsitektur Sistem

│
├── docker-compose.yml # Konfigurasi multi-container
├── Dockerfile # Image untuk aplikasi Go
├── src/
│ ├── go.mod
│ ├── go.sum
│ ├── main.go # Entry point aplikasi
│ ├── views/ # Template HTML (UI)
│ ├── static/ # Asset frontend (CSS, JS, dll)
│ └── ... # File atau folder tambahan
└── mysql_data/ # Volume penyimpanan data MySQL (persistent)


**Container yang dijalankan:**

| Service | Deskripsi | Port |
|----------|------------|------|
| `app` | Aplikasi Golang + Fiber | `3000` |
| `db`  | MySQL Database | `3306` |

---

## Cara Menjalankan

### 1. Pastikan telah terpasang
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

### 2. Jalankan container
```bash
docker compose up --build
