# Golang Testable App

## Install

- Buat Database `resto`

- Copy `.env.example` jadi `.env` dan sesuaikan value-nya **kecuali port aplikasi**, Anda dapat merubah port aplikasi hanya jika Anda benar-benar paham konsekuensinya

- Jalankan `go run seeds/seeder.go` untuk menambahkan menu awal

- Jalakan `docker-compose build && docker-compose up`

## Testing

- Jalankan `go test -coverprofile /tmp/coverage ./... -v`

## Dokumentasi

- Dokumentasi (swagger) dapat diakses melalui `http://localhost:3000/docs/index.html`
