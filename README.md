# flight-booking-api
api for booking flight tickets


### Database

Command to run container for project's database:

 ```
 docker network create my_network
 docker run --name postgre --network my_network \
  -e POSTGRES_PASSWORD=12345 -p 5432:5432 -d postgres
 ```

### Migration
To work on with PostgreSQL database, make sure to migrate up, change DB_URL connection variable in makefile as your postgres configuration and then run:
```
make migrate-up
```


### Redis
Create docker container for redis with following command:

``` docker run -d --name redis -p 6379:6379 redis redis-server --requirepass "12345" ```

### Configs

Do not forget to set all needed configuration variables, for example:

```
export JWT_TOKEN_SECRET=my_secret_key
export REST_PORT=8080
export ACCESS_TOKEN_EXPIRE=877
export ADDRESS=0.0.0.0
export HEADER_TIMEOUT=5s

export POSTGRES_HOST=localhost
export POSTGRES_PORT=5432
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=12345
export POSTGRES_DB=postgres
export POSTGRES_TIMEOUT=30
export POSTGRES_MAX_CONNECTIONS=20

export REDIS_HOST=localhost
export REDIS_PORT=6379
export REDIS_PASSWORD=12345
export REDIS_TIMEOUT=10
export REDIS_TTL=5
export REDIS_DATABASE=0
export REDIS_POOL_SIZE=10
```

After initializing all the necessary dependencies, you can run project:
 ```
 go run cmd/main.go
 ```