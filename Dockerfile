FROM golang:1.16-alpine as builder
WORKDIR /demo/

COPY go.* ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 go build -o /demo/myapp .

FROM alpine:latest
COPY --from=builder /demo/myapp /demo/myapp
EXPOSE 8080
ENTRYPOINT [ "/demo/myapp" ]