FROM golang:1.23 AS build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go vet -v
RUN go generate ./...
RUN go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /
EXPOSE 8081
CMD ["/app", "8081"]
