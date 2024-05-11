include .env
export

LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)

gen:
	@protoc \
			--proto_path=protobuf "protobuf/gallery.proto"\
			--go_out=./protobuf --go_opt=paths=source_relative\
			--go-grpc_out=./protobuf --go-grpc_opt=paths=source_relative
.PHONY: gen

compose-up:
	docker-compose up -d db && docker run mariadb

.PHONY: compose-up

compose-down:
	docker-compose down

.PHONY: compose-down

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'gallery'
.PHONY: migrate-create

run:
	go run cmd/app/main.go

.PHONY: run

mock: 
	mockgen -source ./internal/usecase/interfaces.go -package usecase_test > ./internal/usecase/mocks_test.go

.PHONY: mock