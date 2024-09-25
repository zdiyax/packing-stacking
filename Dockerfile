FROM golang:1.20.5-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM golang:1.20.5-alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app .

FROM scratch
COPY --from=builder /app/docs /docs
COPY --from=builder /app/public /public
COPY --from=builder /bin/app /app
CMD ["/app"]