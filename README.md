# Golang Testable App

## Install

- Buat Database `resto`

- Copy `.env.example` jadi `.env` dan sesuaikan value-nya

- Jalakan `docker-compose build && docker-compose up`

- Jalankan `docker-compose exec app sh -c "go run seeds/seeder.go"`

## Testing

- Jalankan `docker-compose exec app sh -c "go test -coverprofile /tmp/coverage ./... -v"`
