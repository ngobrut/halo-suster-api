migrate-up:
	migrate -path database/migration/ -database "postgresql://postgres:root@localhost:5432/susdb?sslmode=disable" -verbose up

migrate-down:
	migrate -path database/migration/ -database "postgresql://postgres:root@localhost:5432/susdb?sslmode=disable" -verbose down