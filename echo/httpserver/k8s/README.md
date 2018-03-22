# Usage

Prerequisite

* local docker ce channel edge with k8s enabled
* local docker image 'jimlintw/go-app' built

```
kubectl run --image jimlintw/go-app --image-pull-policy=IfNotPresent --port 8080 my-go-app
kubectl expose deploy my-go-app --target-port=8080 --type=LoadBalancer
curl localhost:8080/echo
```
