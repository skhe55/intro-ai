include dev.env
ENV = local
MIGRATIONS-DIR = './migrations/'

ifeq ($(ENV), local)
	PSQL-DSN = "host=localhost user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} port=5433 dbname=${POSTGRES_DB} sslmode=disable"
# Docker build with necessary config params
docker-build: 
	docker-compose --file='./docker/docker-compose.dev.yaml' up --build -d$\ 
	docker exec -it vsftpd usermod -u 1000 ftp$\
	docker exec -it vsftpd chown -R ftp:ftp home/vsftpd/$\
	docker exec -it server-intro make ENV="docker" m-up
# Launch server application
run-server:
	go run ./cmd/main.go
run-server-dev:
	air server --port 3001
else
	PSQL-DSN = "host=db user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} port=5432 dbname=${POSTGRES_DB} sslmode=disable"
endif

m-up:
	goose -dir=${MIGRATIONS-DIR} postgres $(PSQL-DSN) up 
m-down:
	goose -dir=${MIGRATIONS-DIR} postgres ${PSQL-DSN} down
m-status:
	goose -dir=${MIGRATIONS-DIR} postgres $(PSQL-DSN) status
m-reset:
	goose -dir=${MIGRATIONS-DIR} postgres $(PSQL-DSN) reset

# Launch client application in dev mode (by default port is 3001)
run-client:
	cd client && npm run dev