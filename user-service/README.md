---

### Detail Flow Hexagonal Pattern di User Service

- **Model & Entity**
	- `internal/core/domain/model/` berisi struktur data yang merepresentasikan tabel di database (misal: `User`, `Role`, `UserRole`).
	- `internal/core/domain/entity/` berisi representasi data yang digunakan di business logic (misal: `UserEntity`).
	- Model digunakan untuk operasi database, sedangkan entity digunakan untuk proses di service.

- **Repository**
	- `internal/adapter/repository/` berisi interface dan implementasi repository.
	- Interface mendefinisikan kontrak (misal: `GetUserByEmail`).
	- Implementasi repository mengakses database menggunakan model, lalu mengubah hasilnya ke entity.

- **Service/Core**
	- `internal/core/service/` berisi logika bisnis utama, seperti autentikasi user.
	- Service hanya bergantung pada interface repository, tidak pada detail database.
	- Service menerima entity dari repository, memproses, dan mengembalikan hasil ke handler.

- **Handler/Adapter**
	- `internal/adapter/handler/` berisi handler untuk HTTP (Echo).
	- Handler menerima request, validasi, lalu memanggil service.
	- Handler mengubah hasil dari service menjadi response untuk client.

- **Alur Singkat**
	1. Client mengirim request ke endpoint (handler).
	2. Handler validasi dan konversi request ke entity.
	3. Handler panggil service (business logic).
	4. Service panggil repository (akses data).
	5. Repository ambil data dari database (pakai model), ubah ke entity, kembalikan ke service.
	6. Service proses, kembalikan hasil ke handler.
	7. Handler kirim response ke client.

**Kesimpulan:**
Dengan pola ini, setiap lapisan punya tanggung jawab jelas dan bisa diganti tanpa mengubah lapisan lain. Misal, jika ingin ganti database, cukup ubah implementasi repository saja.

---

### Seeder

Seeder digunakan untuk mengisi data awal pada database, seperti data role dan admin. File seeder ada di folder `database/seeds/`:
- `role_seed.go` untuk data role
- `admin_seed.go` untuk data admin

Seeder akan otomatis dijalankan saat koneksi database berhasil (lihat di `config/database.go`).

---

### Cara Menjalankan Project

1. **Clone repository dan masuk ke folder user-service**
2. **Jalankan database dengan Docker Compose:**
	 ```bash
	 docker-compose up -d
	 ```
3. **Jalankan migrasi database:**
	 ```bash
	 migrate -database "postgres://postgres:lokal@localhost:5432/sayur-user-service?sslmode=disable" -path database/migrations up
	 ```
4. **Jalankan aplikasi:**
	 ```bash
	 go run main.go
	 ```
5. **Tes endpoint (misal cek API):**
	 ```http
	 GET http://localhost:8080/api/check
	 ```

Seeder akan otomatis berjalan saat aplikasi pertama kali terhubung ke database.

***

### Kisi-Kisi Implementasi Hexagonal Pattern

* **Repository**: Berisi interface yang mendefinisikan kontrak akses data (misal: ambil, simpan, hapus user). Interface ini harus diimplementasi oleh adapter, seperti repository untuk database PostgreSQL.
* **Service/Core**: Berisi logika bisnis utama, hanya bergantung pada interface, bukan detail teknis.
* **Adapter**: Mengimplementasikan interface dari repository, misal adapter untuk database, HTTP handler, atau pesan dari luar.
* **Handler**: Biasanya berfungsi sebagai penghubung antara dunia luar (API, CLI, dll) dengan service/core.

## Dengan pola ini, Anda bisa mengganti database, protokol, atau teknologi lain tanpa mengubah logika bisnis utama.

### Pola Hexagonal (Ports and Adapters)

User Service ini menggunakan pola arsitektur Hexagonal (juga dikenal sebagai Ports and Adapters). Pola ini bertujuan agar kode inti aplikasi (bisnis logic) terpisah dari detail teknis seperti database, HTTP, atau pesan dari luar. Dengan pola ini, aplikasi lebih mudah untuk diuji, dikembangkan, dan diintegrasikan dengan teknologi lain.

**Kesimpulan singkat:**
Hexagonal Pattern membuat aplikasi lebih fleksibel dan mudah beradaptasi dengan perubahan, karena bagian inti tidak bergantung langsung pada teknologi eksternal.

## Dokumentasi User Service (Bahasa Indonesia)

> **Catatan:** Dokumentasi ini akan diupdate secara berkala. Silakan tambahkan atau ubah bagian ini sesuai kebutuhan. Dokumentasi ini tidak mengganggu tulisan Anda yang sudah ada di bawah.

### Deskripsi Singkat

User Service adalah layanan yang bertanggung jawab untuk mengelola data pengguna, autentikasi, otorisasi, dan peran (role) dalam aplikasi micro-sayur. Service ini menangani pendaftaran, login, manajemen user, serta pengelolaan role dan relasi user-role.

### Fitur Utama

* Registrasi dan login pengguna
* Manajemen data user (CRUD)
* Manajemen role dan user-role
* Integrasi dengan database PostgreSQL
* Migrasi database menggunakan Go Migrate
* Seed data awal (admin dan role)

### Struktur Folder Penting

* `cmd/` : Entry point aplikasi
* `config/` : Konfigurasi aplikasi (database, redis, rabbitmq)
* `database/migrations/` : File migrasi database
* `database/seeds/` : File seed data awal
* `internal/core/domain/model/` : Model data user, role, user-role
* `internal/adapter/handler/` : Handler untuk HTTP/gRPC
* `internal/adapter/repository/` : Repository untuk akses data

### Cara Menjalankan

Lihat bagian di bawah untuk instruksi menjalankan migrasi dan perintah lainnya.




Contoh cara jalankan Migrate


### Catatan & Instruksi Migrasi

**Menjalankan Migrasi Database:**

```bash
migrate -database "postgres://postgres:lokal@localhost:5432/sayur-user-service?sslmode=disable" -path database/migrations up
```

**Jika Go Migrate gagal:**

```bash
go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

**Jika terdapat Dirty database version:**

```bash
migrate -database "postgres://postgres:lokal@localhost:5432/sayur-user-service?sslmode=disable" -path database/migrations force <version>
```

Ganti `<version>` dengan versi yang ingin di-force.

**Cara pakai psql dengan nama database:**

```bash
docker exec -it <container_id> psql -U postgres -d sayur-user-service
```

Ganti `<container_id>` dengan ID container PostgreSQL Anda.


