FROM golang:latest as builder
RUN mkdir /app
COPY soma_test.go /app/
WORKDIR /app
RUN go build -o main .


FROM scratch

COPY --from=builder /app /app
CMD ["/app/main"]

