FROM alpinelinux/golang:latest AS builder

WORKDIR /usr/src/app

COPY . .

RUN go build -o full-cycle-rocks

FROM scratch

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/full-cycle-rocks .

ENTRYPOINT [ "./full-cycle-rocks" ]