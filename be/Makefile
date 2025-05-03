# Thông tin SSH (máy chủ từ xa)
SSH_USER=server0103
SSH_HOST=192.168.1.22
SSH_PORT=22

# Thông tin PostgreSQL (trên máy chủ từ xa)
DB_NAME=simple_bank
DB_USER=postgres
DB_HOST=192.168.1.22
DB_PORT=5432

# Lệnh psql qua SSH
SSH_CMD=ssh -p $(SSH_PORT) $(SSH_USER)@$(SSH_HOST)
PSQL_CMD=PGPASSWORD="123" psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER)


# Tạo database từ xa
create_db:
	$(SSH_CMD) '$(PSQL_CMD) -c "CREATE DATABASE $(DB_NAME);"'

# Xóa database từ xa
drop_db:
	$(SSH_CMD) '$(PSQL_CMD) -c "DROP DATABASE IF EXISTS $(DB_NAME);"'

migrate_up:
	migrate -path db/migration -database "postgresql://postgres:123@192.168.1.22:5432/simple_bank" -verbose up

migrate_up1:
	migrate -path db/migration -database "postgresql://postgres:123@192.168.1.22:5432/simple_bank" -verbose 1

migrate_down:
	migrate -path db/migration -database "postgresql://postgres:123@192.168.1.22:5432/simple_bank" -verbose down

migrate_down1:
	migrate -path db/migration -database "postgresql://postgres:123@192.168.1.22:5432/simple_bank" -verbose 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go  github.com/phuongtnd/test01/db/sqlc Store

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: drop_db create_db migrate_up migrate_down migrate_up1 migrate_down1 redis
 


