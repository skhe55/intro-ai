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

#run docker-container
dftp: 
	docker-compose --file='./docker/docker-compose.dev.yaml' up --build -d && docker exec -it vsftpd usermod -u 1000 ftp && docker exec -it vsftpd chown -R ftp:ftp home/vsftpd/

#run app
start:
	go run ./cmd/main.go
dev:
	air server --port 3001