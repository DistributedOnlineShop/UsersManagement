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
	rm -f pb/session/*.go
	rm -f pb/addresses/*.go

	protoc --proto_path=proto --go_out=pb/users --go_opt=paths=source_relative \
	--go-grpc_out=pb/users --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb/users --grpc-gateway_opt paths=source_relative \
    proto/service_users.proto

	protoc --proto_path=proto --go_out=pb/session --go_opt=paths=source_relative \
	--go-grpc_out=pb/session --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb/session --grpc-gateway_opt paths=source_relative \
    proto/service_session.proto

	protoc --proto_path=proto --go_out=pb/addresses --go_opt=paths=source_relative \
	--go-grpc_out=pb/addresses --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb/addresses --grpc-gateway_opt paths=source_relative \
    proto/service_addresses.proto

.PHONY: postgres dropPsql createDB dropDB migrateup migratedown sqlc test server protob