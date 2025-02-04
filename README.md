# GO RESTful API with Postgressql

## Description
Build a RESTful API for **monitoring money** using Golang and implement PostgreSQL as the database.

## Configuration
All configuration is in `.env.example` file, save as `.env`
```
DB_USERNAME=
DB_PASSWORD=
DB_HOST=localhost
DB_PORT=5432
DB_NAME=
DB_MAX_IDLE_CONNS=10
DB_MAX_OPEN_CONNS=50
DB_CONN_MAX_IDLE=5
DB_CONN_MAX_LIFE=60

SERVER_ADDR=localhost:8080 //you can setup port server
```

## DB Migration
You can copy sql table in `db.sql` to your local db.