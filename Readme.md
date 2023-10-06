# Tabungan API

## Description

API to store user and transaction bank account


## Installation
### Migrate database
```
docker compose up -d \
make migrate url="postgres://{user_db}:{password_db}@{host_db}:{port_db}/{db_name}?sslmode=disable"
```
### Download depedencies
make vendor