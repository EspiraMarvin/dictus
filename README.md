## golang crud with GORM and POSTGRES

1. Create an .env file and populate the variables below.

```bash

    # APP PORT
    PORT=

    # JWT
    SECRET=""

    # PostgreSQL
    DB_USER=""
    DB_PASS=""
    DB_HOST=""
    DB_PORT=""
    DB_NAME=""

```

2.  Alternatively, you could create an `env.sh` and add the following environment variables. This assumes that you have
    created a database whose information is populated under `DB_` prefix.

```bash

   # APP PORT
   export PORT=""

   # JWT
   export SECRET=""

   # PostgreSQL
   export DB_USER=""
   export DB_PASS=""
   export DB_HOST=""
   export DB_PORT=""
   export DB_NAME=""
```
