FROM golang:1.13.4-alpine3.10 AS builder

ARG APPPATH=$GOPATH/src/github.com/jonatascabral/jokes-app

COPY . $APPPATH

WORKDIR $APPPATH

RUN apk --no-cache add tzdata dep git

RUN dep ensure -v && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -o /app ./cmd

FROM alpine:3.10

ARG APPPATH=/go/src/github.com/jonatascabral/jokes-app
ARG TZ=America/Sao_Paulo

WORKDIR /app/

COPY --from=builder /app ./
COPY --from=builder /usr/share/zoneinfo/$TZ /etc/localtime
COPY --from=builder $APPPATH/.env ./
COPY --from=builder $APPPATH/resources ./

RUN echo $TZ > /etc/timezone && \
    chmod +x /app

ENTRYPOINT ["./app"]

EXPOSE 8080