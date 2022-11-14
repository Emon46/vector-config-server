# vector config api server
this api server will help to get and update `configMap`.

## get vector configmap
Sample `GET` Request:
``` 
127.0.0.1:8080/transforms
```
Sample Request Body:
``` 
{
    "configMapName": "vector-data-plane-config",
    "configMapNameSpace": "vector",
}
```
## update vector configmap
Sample `POST` Request:
``` 
127.0.0.1:8080/transforms
```
Sample Request Body:
``` 
{
    "configMapName": "vector-data-plane-config",
    "configMapNameSpace": "vector",
    "transforms": {
        "fillter_test_config": {
            "condition": "contains(string(.message) ?? \"\", \"no_tag\") != true",
            "inputs": [
            "k8s_logs_source"
            ],
            "type": "filter"
        }

    }
}
```

here we have added a new `transform` to our `vector configMap`.  

## run unit test
``` 
make test
```
This command will run against all the unit test.

## build Binary
``` 
make build
```

This command will build the binary inside `bin/` repo.

## docker build

``` 
export REGISTRY=<docker-hub-Registry-name>
make docker-build
```
this command will build a docker image.

## docker push

``` 
export REGISTRY=<docker-hub-Registry-name>
make docker-push
```
this command will push a docker image in your `docker-hub` repository.

## load to Kind cluster

``` 
export REGISTRY=<docker-hub-Registry-name>
make push-to-kind
```
this command will build and load the docker image in your local kind cluster.