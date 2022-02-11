# Go pgx Example

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](/LICENSE)

The Golang [jackc/pgx](https://github.com/jackc/pgx) PostgreSQL usage example.

## Requirement

1. [PostgreSQL](https://www.postgresql.org/download/)
2. Setup environment variable:
   1. `GOPGX_DATABASE_URL`, the PostgreSQL database connection url. eg. `postgres://username:password@localhost:5432/gopgx_example`

### Database migration

In order to have runable for each example of this repo, you need to execute the provided [database.sql](/scripts/database.sql) script beforehand.

Example of execution using `psql`:

```shell
PGPASSWORD=password psql -U username -d gopgx_example -h localhost -p 5432 < scripts/database.sql
```

## Example

1. [hello world](/cmd/helloworld/main.go), the very simple example of hello world query using `pgx`.
2. [hello world pool](/cmd/helloworldpool/main.go), the example of config the pgx connection pooling using `pgxpool`, then a simple hello world query.
3. [member model](/internal/models/model.go), the example of mapping the postgresql table [structure](/scripts/database.sql) to Go struct.

## Copyright and License

Copyright 2022 Maikel Chandika (mkdika@gmail.com). Code released under the MIT License. See [LICENSE](/LICENSE) file.