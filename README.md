# Golang Testable App

## Install

- Buat Database `resto`

- Copy `.env.example` jadi `.env` dan sesuaikan value-nya **kecuali port aplikasi**, Anda dapat merubah port aplikasi hanya jika Anda benar-benar paham konsekuensinya

- Jalakan `docker-compose build && docker-compose up`

- Jalankan `docker-compose exec app sh -c "go run seeds/seeder.go"` untuk menambahkan menu awal

## Testing

- Jalankan `docker-compose exec app sh -c "go test -coverprofile /tmp/coverage ./... -v"`

## Dokumentasi

- Dokumentasi dapat diakses melalui `http://localhost:3000/docs/index.html`
