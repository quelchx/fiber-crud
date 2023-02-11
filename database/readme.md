# Database Directory

This directory is setup to use Prisma to simply pull the current database schema and connect to the database in the web browser.

This will run the a GUI for the databse on `localhost:5555`

## Setup

1. Install Prisma CLI

```bash
npm install -g prisma
```

2. Install dependencies

```bash
npm install
```

3. Start the database

```bash
npm run start
```

## Manual Setup

This assumes you have some connection to a database (container or external)

```bash
mkdir database && cd database
yarn init -y

yarn add -D prisma

mkdir prisma && cd prisma
touch schema.prisma
```

## Prisma Schema

```prisma
datasource db {
  provider = "postgresql"
  url      = env("DATABASE_URL")
}

generator client {
  provider = "prisma-client-js"
}

```

## Environment Variables

```bash
touch .env
```

```env
DATABASE_URL="postgresql://user:password@localhost:5432/database?schema=public"
```

## Start the Database

```bash
prisma db pull
prisma studio
```

## Resources

- [Prisma Docs](https://www.prisma.io/docs/)
