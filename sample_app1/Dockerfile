FROM golang:1.23rc2-alpine3.20 as build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:3.20.1 as golang

COPY --from=build /src/main /src/main

CMD ["/src/main"]