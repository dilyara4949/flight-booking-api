# flight-booking-api
api for booking flight tickets


### Testing

to test integration tests:
```
make testintegration
```
Before testing migrate up testing database by command:
```
make migrate-up-test
```

also set config variables as below example:

```
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=12345
POSTGRES_DB=postgrestest
POSTGRES_TIMEOUT=30
POSTGRES_MAX_CONNECTIONS=20

JWT_TOKEN_SECRET=my_secret_key
REST_PORT=8080
ACCESS_TOKEN_EXPIRE=877
ADDRESS=0.0.0.0
```