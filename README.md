#### start environment
```shell
docker-compose up -d
```

#### access to container
```shell
docker exec -it appproduct bash
```

#### generate mocks with mockgen
```shell
mockgen -destination=application/mocks/application.go -source=application/product.go application
```

#### testing app

```shell
go test ./...
```