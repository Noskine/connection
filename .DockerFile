FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM scratch

COPY --from=builder /app/main /main

ENV DB_USER=admin \
    DB_PASSWORD=T5urIIcFZ47nyuvKCkk8JCnA0WzmA216EHFYPTf0nPCqJQsJvU \
    DB_NAME=testdb \
    DB_HOST=db

ENTRYPOINT ["/main"]
