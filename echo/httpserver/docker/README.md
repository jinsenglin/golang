# Build Minimal Image

```
docker build -t jimlintw/go-app .
```

Result 6.83 MB

# Use the Minimal Image

```
docker run --rm -d -p 8080:8080 jimlintw/go-app go-app
```

# Reference

* https://blog.wu-boy.com/2017/04/build-minimal-docker-container-using-multi-stage-for-go-app/
