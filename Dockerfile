FROM golang:latest as builder
RUN mkdir /app
COPY *.go /app/
WORKDIR /app
RUN go build -o soma .


FROM scratch

COPY --from=builder /app /app
CMD ["/app/soma"]

