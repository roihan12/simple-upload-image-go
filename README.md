# simple-upload-image-go



# Endpoint API

## Endpoint Auth

Endpoint ini bertanggung jawab untuk login dan register.
Method | Path | Keterangan | Auth
------------- | ------------- | ------------- | -------------
**_POST_** | _`/api/v1/users/login`_ | Untuk masuk atau login ke akun yang telah dibuat .
**_POST_** | _`/api/v1/users/register`_ | Membuat akun user.

## Endpoint User

Endpoint ini bertanggung jawab update dan delete user.
Method | Path | Keterangan | Auth
------------- | ------------- | ------------- | -------------
**_PUT_** | _`/api/v1/users/:id`_ | Mengubah data pengguna berdasakan id pengguna | token
**_DELETE_** | _`/api/v1/users/:id`_ | Menghapus data pengguna berdasakan id pengguna | token

## Endpoint Photo

Endpoint ini bertanggung jawab mengelola photo.
Method | Path | Keterangan | Auth
------------- | ------------- | ------------- | -------------
**_GET_** | _`/api/photos`_ | Mengakses data foto berdasarkan user id | token
**_GET_** | _`/api/photos/:id`_ | Mengakses data pengguna berdasakan id foto | token
**_POST_** | _`/api/photos`_ | Membuat data foto baru | token
**_PUT_** | _`/api/photos/:id`_ | Mengubah data foto berdasakan id foto | token
**_DELETE_** | _`/api/photos/:id`_ | Menghapus data foto berdasakan id foto | token

## Instalasi

Untuk menjalankan proyek ini secara lokal, pastikan Anda telah menginstal GoLang. Berikut langkah-langkah umum instalasi GoLang:

1. **Unduh GoLang**: Unduh dan instal GoLang dari [situs resmi GoLang](https://golang.org/dl/).

2. **Konfigurasi Lingkungan**: Setelah menginstal, pastikan Anda mengatur variabel lingkungan seperti `GOPATH` dan `PATH` sesuai dengan dokumentasi GoLang.

3. **Klon Proyek**: Klon repositori ini ke komputer Anda.
4. **Buka Postgres** Buat database baru kemudian copy file .env_example
5. **Jalankan**: Buka terminal, masuk ke direktori proyek, dan jalankan aplikasi dengan perintah:

   ```bash
   # 1. Buka terminal.

   # 2. Pastikan Anda berada dalam direktori proyek.

   # 3. Jalankan aplikasi dengan perintah berikut:
   go run main.go
   ```

   Untuk menjalankan proyek ini menggunakkan docker, pastikan Anda telah menginstal Docker.

**Jalankan**: Buka terminal, masuk ke direktori proyek, dan jalankan aplikasi dengan perintah:

```bash
# 1. Buka terminal.

# 2. Pastikan Anda berada dalam direktori proyek.

# 4. Jika menggunakan docker. Jalankan aplikasi dengan perintah berikut:
   - docker compose build
	- docker compose -f docker-compose.yaml --profile tools run --rm migrate up
	- docker compose up
# 3. Jika anda bisa menggunakan makefile. Jalankan aplikasi dengan perintah berikut:
   - make build
   - make migrate-up
```

Aplikasi akan berjalan di http://localhost:3000

## Tools 🛠

[![My Skills](https://skillicons.dev/icons?i=go,git,postgres,postman,vscode,docker)](https://skillicons.dev)
