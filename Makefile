always-run:

vendor: always-run
	env GO111MODULE=on go mod vendor

migrate:
	# https://github.com/golang-migrate/migrate
	migrate -path migrations -database $(url) up
