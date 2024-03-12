# Build Stage

FROM golang:alpine AS build

WORKDIR /temp
COPY . .
RUN go build -o build/bin -v ./cmd


# Deploy Stage

FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app
COPY --from=build /temp/build/bin /app/bin

CMD ["/app/bin", "server"]