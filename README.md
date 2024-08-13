# flight-booking-api
api for booking flight tickets

### Configuration

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

### Kubernetes

[Install minikube if you dont have it yet](https://minikube.sigs.k8s.io/docs/start/)

The you just need too, apply kubernetes configuration to run the project:

```
make apply-kube
```

Warning: job 'migrate-job' may crash before 'app' pod is running, that is why you may recreate  'migrate-job' to run migration on postgres schema after 'app' pod is ready.

To forward a local port 8080 to a port on a Kubernetes service:
```
make kube-forward-port
```