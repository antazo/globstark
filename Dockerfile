# Backend Dockerfile
FROM golang:latest AS builder

WORKDIR /globstark-app

COPY go.mod ./
# COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /globstark ./cmd/server

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /globstark .
COPY --from=builder /globstark-app/web ./web

EXPOSE 7979

CMD ["./globstark"]
#CMD ["ls", "-la"]