FROM golang:1.12 as build
WORKDIR /app
COPY . .
RUN GOPROXY=https://goproxy.io CGO_ENABLED=0 go build -o base-go cmd/main.go

FROM alpine:3.4 as run
COPY --from=build /app/base-go /app/base-go
ENTRYPOINT [ "/app/base-go" ]
EXPOSE 5000
