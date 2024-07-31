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
make export_env
source set_env.sh
```