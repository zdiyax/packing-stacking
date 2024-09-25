FROM golang:1.20.5-alpine as modules
COPY go.mod go.sum /modules/
COPY vendor /modules/vendor
WORKDIR /modules

# Step 2: Builder
FROM golang:1.20.5-alpine as builder
COPY --from=modules /modules/vendor /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app .

# Step 3: Final
FROM scratch
COPY --from=builder /app/docs /docs
COPY --from=builder /app/public /public
COPY --from=builder /bin/app /app
CMD ["/app"]