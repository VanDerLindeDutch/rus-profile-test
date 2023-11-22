FROM golang:1.21 as build-env

WORKDIR /usr/local/go/src/rus-profile-test

COPY . .

RUN go get ./...
RUN go build -o /rus-profile-test rus-profile-test/cmd/grpc_server

FROM ubuntu

COPY --from=build-env /rus-profile-test /
RUN mkdir -p /api
COPY --from=build-env /usr/local/go/src/rus-profile-test/api /api



CMD ["/rus-profile-test"]

EXPOSE 8081
