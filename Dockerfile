FROM golang:latest as builder
LABEL authors="antoq"
WORKDIR /build
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o ./main


FROM scratch
WORKDIR /app
COPY --from=builder /build/main ./main
COPY --from=builder /build/templates ./templates
EXPOSE 8080
ENTRYPOINT ["./main"]