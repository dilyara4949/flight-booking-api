# flight-booking-api
api for booking flight tickets

### Database

Command to run container for project's database:

 ```
 docker run --name postgres --network network_name /
  -e POSTGRES_PASSWORD=12345 -p 5432:5432 -d postgres
 ```

To run this command, you will need official postgres image.

### Migration
To work on with PostgreSQL database, make sure to migrate up, change DB_URL connection variable in makefile as your postgres configuration and then run:
```
make migrate-up
```

### Configs

Do not forget to set all needed configuration variables, for example: 

```
export POSTGRES_HOST=localhost
export POSTGRES_PORT=5432
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=12345
export POSTGRES_DB=postgres
export POSTGRES_TIMEOUT=30
export POSTGRES_MAX_CONNECTIONS=20

export JWT_TOKEN_SECRET=my_secret_key
export REST_PORT=8080
export ACCESS_TOKEN_EXPIRE=877
export ADDRESS=0.0.0.0
```

After initializing all the necessary dependencies, you can run project:
 ```
 go run cmd/main.go
 ```