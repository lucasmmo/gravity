FROM golang:1.19.1 as builder
WORKDIR /build
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app cmd/api/*.go


FROM alpine:latest
EXPOSE 8080
WORKDIR /app
COPY --from=builder /build/app .
CMD [ "./app" ]