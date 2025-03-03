DBName = UsersManagement
User = root
Password = secret

createDB:
	docker exec -it psql createdb --username=root --owner=root $(DBName)

dropDB:
	docker exec -it psql dropdb $(DBName)

# Migrate 初期設定
# createMigrate:
# 	migrate create -ext sql -dir db/migration -seq init_table

#up
migrateup:
	migrate -path db/migration -database postgresql://${User}:${Password}@localhost:5432/${DBName}?sslmode=disable -verbose up

#down
migratedown:
	migrate -path db/migration -database postgresql://${User}:${Password}@localhost:5432/${DBName}?sslmode=disable -verbose down

#Sqlc
sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run .

protob:
	rm -f pb/users/*.go
	rm -f pb/verification/*.go

	protoc --proto_path=proto --go_out=pb/users --go_opt=paths=source_relative \
	--go-grpc_out=pb/users --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb/users --grpc-gateway_opt paths=source_relative \
    proto/service_users.proto

	protoc --proto_path=proto --go_out=pb/verification --go_opt=paths=source_relative \
	--go-grpc_out=pb/verification --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb/verification --grpc-gateway_opt paths=source_relative \
    proto/service_verification.proto

.PHONY: postgres dropPsql createDB dropDB migrateup migratedown sqlc redis dropRedis test server protob