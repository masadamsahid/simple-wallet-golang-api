To generate migration files:

```bash
migrate create -ext=sql -seq -dir="./db/migrations" FILE_NAME
```

To run

```bash
migrate -path="./db/migrations" -database="postgres://postgres:postgres@localhost:5432/mampuio_wallet_db?sslmode=disable" up
```