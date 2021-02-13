FROM golang:alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /builder
COPY . .
RUN go mod download
RUN go build -o app main.go

FROM scratch

COPY --from=builder /build/app .
ENTRYPOINT [ "./app" ]
EXPOSE 3000