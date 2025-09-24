FROM golang:1.21-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /usr/local/bin/auth-service ./cmd/auth-service

FROM alpine:3.18
COPY --from=build /usr/local/bin/auth-service /usr/local/bin/auth-service
EXPOSE 8080
CMD ["/usr/local/bin/auth-service"]