# packing-stacking

[Demo](https://packing-stacking-b591958da039.herokuapp.com/) for you to enjoy.

### How to run it

1. You can run it locally: 
```shell
PORT=1234 go run main.go
```

2. You can run the container:
```shell
docker build . -t packnstack && docker run -e PORT='1234' -p 1234:1234 -t packnstack
```

You'll find it to be available on `localhost:1234` address in your browser.


### Docs

You can also access Swagger documentation here:

```
localhost:1234/swagger/
```