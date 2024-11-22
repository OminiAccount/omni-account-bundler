FROM golang:1.22.4 AS build
RUN mkdir /src
WORKDIR /src
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /src/oa-bundler /src/cmd

FROM ubuntu:latest
RUN apt-get update
RUN apt-get install -y ca-certificates
RUN update-ca-certificates
RUN mkdir /app
COPY --from=build /src/oa-bundler /app/oa-bundler

WORKDIR /app
EXPOSE 8100
CMD ["/bin/sh", "-c", "/app/oa-bundler"]
