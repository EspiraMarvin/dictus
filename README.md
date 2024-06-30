## golang crud with GORM and POSTGRES

1 . Create `env.sh` and add the following environment variables. This assumes that you have
created a database whose information is populated under `DB_` prefix.
```bash

    # App post
    export PORT=""

    # PostgreSQL
    export DB_USER=""
    export DB_PASS=""
    export DB_HOST=""
    export DB_PORT=""
    export DB_NAME=""
    ```

2. Alternatively create an .env file and populate the variables

   ```bash
   # APP PORT
   PORT=

   # PostgreSQL
   DB_USER=""
   DB_PASS=""
   DB_HOST=""
   DB_PORT=""
   DB_NAME=""

   ```
