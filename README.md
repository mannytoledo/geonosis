# geonosis

One click app deploy

# Run with godep

```bash
cd geonosis
godep go run geonosis/main.go
```

# Configuration

Currently, geonosis will use the same environment variables as `docker-machine`

Env var  | example value
------------- | -------------
DOCKER_HOST  | tcp://192.168.99.100:2376
DOCKER_CERT_PATH  | /Users/you/.docker/machine/machines/dev
