# Comando para criar uma nova migration
migrate-create: 
	migrate create -ext=sql -dir=database/migrations -seq $(name)
run:
	docker compose up -d 
	go run cmd/main.go

swagger:
	swag init -g cmd/main.go