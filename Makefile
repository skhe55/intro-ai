include dev.env
# goose migrations
PSQL-DSN = "host=localhost user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} port=5433 dbname=${POSTGRES_DB} sslmode=disable"
MIGRATIONS-DIR = './migrations/'

m-up:
	goose -dir=${MIGRATIONS-DIR} postgres $(PSQL-DSN) up 
m-down:
	goose -dir=${MIGRATIONS-DIR} postgres ${PSQL-DSN} down
m-status:
	goose -dir=${MIGRATIONS-DIR} postgres $(PSQL-DSN) status

#run app
start:
	go run ./cmd/main.go