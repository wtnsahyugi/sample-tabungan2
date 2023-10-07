always-run:

vendor: always-run
	env GO111MODULE=on go mod vendor

migrate:
	# https://github.com/golang-migrate/migrate
	migrate -path migrations -database $(url) up

migration:
	# https://github.com/golang-migrate/migrate
	migrate create -ext sql -dir migrations -seq $(name)

rollback:
	# https://github.com/golang-migrate/migrate
	migrate -path migrations -database $(url) down 1
