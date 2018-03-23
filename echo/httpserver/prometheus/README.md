# Usage

```
go get -u github.com/prometheus/client_golang/prometheus

go run main.go

curl http://localhose:8080/echo
curl http://localhost:8080/metrics
```

# Build Docker Image

```
go get -u github.com/prometheus/client_golang/prometheus

govendor init
govendor add +external

docker build -t jimlintw/go-app-with-metrics .
```
