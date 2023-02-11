# Go Fiber CRUD API Example

This is a simple CRUD API example using Go Fiber, GORM, and PostgreSQL.

## Table Of Contents

- [Go Fiber CRUD API Example](#go-fiber-crud-api-example)
  - [Table Of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
    - [Packages Used](#packages-used)
  - [To Create Postgres Container](#to-create-postgres-container)
    - [Run Application](#run-application)
  - [Optional](#optional)
    - [Migrations](#migrations)
    - [Database GUI (Prisma)](#database-gui-prisma)

## Prerequisites

- Go 1.16+
- Docker

### Packages Used

- Fiber
- GORM (PostgreSQL)
- GoDotEnv
- CompileDaemon

## To Create Postgres Container

**Make sure to configre your .env file**

```sh
docker-compose up -d
```

### Run Application

```sh
CompileDaemon -command='/fiber-crud'

# or

CompileDaemon -command='go run main.go"
```

## Optional

The below steps are optional and not required to run the application.

### Migrations

**Note:** This is not required as the `initializers` package includes a function to auto migrate the database on startup.

```sh
go run migrate/migrate.go
```

### Database GUI (Prisma)

I know they're many ways to connect to your database and view the data. I chose to use Prisma because it's easy to setup and use. Others may use PGAdmin, DBeaver, etc.

**Note:** This requires `nodejs` and `npm` to be installed on your machine.

View the database `readme.md` for more information.

[Database GUI with Prisma Studio](database/readme.md)

```sh
cd database
npm run start
```

This will spin up a GUI for the database on `localhost:5555`
