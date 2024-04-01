FROM golang:alpine AS blogs-builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o blogs-webapp

FROM alpine
COPY --from=blogs-builder /app/blogs-webapp /usr/bin/blogs-webapp
EXPOSE 8082
ENTRYPOINT ["/usr/bin/blogs-webapp"]