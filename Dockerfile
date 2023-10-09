FROM golang:alpine AS builder
WORKDIR /go/src/github.com/kosuke-taniguchi/k8s-sample
COPY main.go go.mod go.sum ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o k8s-sample .

FROM alpine:latest
COPY --from=builder /go/src/github.com/kosuke-taniguchi/k8s-sample/k8s-sample .
CMD [ "./k8s-sample" ]