FROM golang:1.16-alpine AS build

COPY go.mod /go/src/httpserver
WORKDIR /go/src/httpserver
RUN go mod download

COPY *.go /go/src/httpserver
RUN go build -o /bin/httpserver

FROM alpine
COPY --from=build /bin/httpserver /bin/httpserver
ENTRYPOINT["/bin/httpserver"]