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
source set_env.sh
```

After initializing all the necessary dependencies, you can run project:
 ```
 go run cmd/main.go
 ```

Use, if running on local machine:
``` 
POSTGRES_HOST=localhost
REDIS_HOST=localhost
```
if  running on docker or kubernetes:
```
POSTGRES_HOST=postgres
REDIS_HOST=redis
```

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

### Kubernetes

[Install minikube if you dont have it yet](https://minikube.sigs.k8s.io/docs/start/)

The you just need too, apply kubernetes configuration to run the project:

```
helm install booking-app ./booking-chart
```

or to update:
```
helm upgrade booking-app ./booking-chart
```

Warning: job 'migrate-job' may crash before 'app' pod is running, that is why you may recreate  'migrate-job' to run migration on postgres schema after 'app' pod is ready.

To forward a local port 8080 to a port on a Kubernetes service:
```
make kube-forward-port
```
