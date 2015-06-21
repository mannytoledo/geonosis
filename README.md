# geonosis

### Development setup

NOTE: Assumes you have docker-machine up and running! See configuration section below if you are unsure.

1. Clone geonosis into `/path/to/go/workspace/github.com/behanceops/genosis` with `git clone https://github.com/behanceops/geonosis /path/to/go/workspace/github.com/behanceops/genosis`
2. Install godeps with `go get github.com/tools/godep`
3. cd to `/path/to/go/workspace/github.com/behanceops/genosis/Godeps` and run `godep restore` to get all dependencies described in `Godeps.json`
4. cd to `/path/to/go/workspace/github.com/behanceops/genosis` and run `go run main.go`, then goto `http://localhost:1323/` in your browser.

One click app deploy

### Run with godep

```bash
cd geonosis
godep go run geonosis/main.go
```

### Configuration

Currently, geonosis will use the same environment variables as `docker-machine`

Env var  | example value
------------- | -------------
DOCKER_HOST  | tcp://192.168.99.100:2376
DOCKER_CERT_PATH  | /Users/you/.docker/machine/machines/dev
