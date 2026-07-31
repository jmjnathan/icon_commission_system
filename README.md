# Icon Commission System — Sacred Icon Studio

Aplikasi fullstack untuk manajemen komisi lukisan ikon, mencakup data client, master data, tracking pesanan (commission), serta pencatatan arus kas sederhana.

Dibuat sebagai bagian dari Technical Test — Fullstack Developer.

## Tech Stack

**Backend**

- Go + Gin Framework
- GORM + PostgreSQL
- Clean Architecture (entity, dto, repository, usecase, handler, router)
- JWT Authentication
- golang-migrate (dokumentasi struktur migrasi database)

**Frontend**

- Vue 3 (Composition API) + TypeScript
- Tailwind CSS
- Vue Router
- Axios (dengan interceptor untuk auth token & error handling)

## Fitur Utama

- **Autentikasi** — Login admin menggunakan JWT
- **Master Data** — CRUD untuk Ukuran, Material, Style, dan Santo/Santa (mendukung autocomplete di form Commission)
- **Client Management** — Data lengkap pemesan, termasuk preferensi santo pelindung dan riwayat pesanan
- **Commission (Order Komisi)** — Mendukung banyak item dalam satu transaksi (pola header-detail), lengkap dengan foto referensi, tracking status (pending → in_progress → completed → delivered)
- **Riwayat Transaksi** — Daftar komisi yang sudah selesai/terkirim, dengan fitur cetak label pengiriman
- **Arus Kas Keluar (Cash Out)** — Pencatatan pengeluaran (qty, harga satuan, vendor, metode bayar, bukti struk)
- **Laporan Keuangan** — Ringkasan kas masuk/keluar dan laba-rugi sederhana per bulan, dapat dicetak
- **Upload File** — Foto referensi komisi dan bukti struk pengeluaran, disimpan di server (folder `uploads/`)

## Cara Menjalankan Project

### 1. Backend

```bash
cd backend
```

Buat file `.env`:
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=02102002
DB_NAME=icon_commission_db
APP_PORT=8080
JWT_SECRET=rahasia-super-aman-ganti-nanti

Buat database (lewat psql atau pgAdmin):

```sql
CREATE DATABASE icon_commission_db;
```

Install dependency & jalankan:

```bash
go mod tidy
go run main.go
```

Server berjalan di `http://localhost:8080`. Struktur tabel akan dibuat otomatis oleh GORM AutoMigrate, dan admin default akan otomatis di-seed.

**Kredensial Admin Default:**
Username: admin
Password: admin123

### 2. Frontend

```bash
cd frontend
npm install
npm run dev
```

Aplikasi berjalan di `http://localhost:5173`.

## Dokumentasi API

Koleksi Postman tersedia di file [`docs/Icon-Commission-System.postman_collection.json`](./docs/) — import ke Postman/Insomnia untuk mencoba seluruh endpoint API.

## Catatan Migrasi Database

Struktur database dikelola menggunakan **GORM AutoMigrate** selama tahap development untuk mempercepat iterasi skema. Sebagai dokumentasi resmi struktur akhir database, tersedia juga file migrasi SQL menggunakan **golang-migrate** di folder `backend/migrations_sql/`, berisi definisi `up`/`down` untuk seluruh tabel (admins, master data, clients, commissions, commission_items, commission_photos, cash_outs).

## Struktur Folder Backend

backend/
├── config/ → koneksi database
├── internal/
│ ├── entity/ → model data (per modul: admin, client, master, commission, cashout)
│ ├── dto/ → request/response struct + validasi
│ ├── repository/ → akses database (GORM)
│ ├── usecase/ → business logic
│ ├── handler/ → controller Gin
│ ├── middleware/ → JWT auth middleware
│ └── router/ → definisi route per modul
├── migrations_sql/ → file migrasi SQL (golang-migrate)
├── uploads/ → penyimpanan file upload
└── main.go

## Struktur Folder Frontend

frontend/src/
├── views/ → halaman per fitur (dashboard, commission, client, master, cashflow)
├── components/ → komponen reusable (table, modal, toast, confirm, input, upload)
├── composables/ → logic reusable (auth, commission, client, master, cashout)
├── services/ → axios instance + interceptor
└── router/ → konfigurasi vue-router
