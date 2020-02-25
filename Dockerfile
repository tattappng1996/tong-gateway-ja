FROM golang:1.13.7 AS builder
WORKDIR /go/src/tong-gateway-ja
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /tong-gateway-ja .

FROM scratch
COPY --from=builder ./tong-gateway-ja ./
ENTRYPOINT ["/tong-gateway-ja"]