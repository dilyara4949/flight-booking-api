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

Do not forget to set all needed configuration variables:

```
make export_env
source export_env
```

After initializing all the necessary dependencies, you can run project:
 ```
 go run cmd/main.go
 ```