package gen

//go:generate docker compose up -d
//go:generate go run ./cmd/tools/terndotenv/main.go
//go:generate sqlc generate -f ./internal/store/pgstore/sqlc.yml
//go:generate cyclonedx-gomod mod -licenses -type application -output bom.xml
//go:generate govulncheck ./...
//go:generate go run ./cmd/gows/main.go
