#
# Builder
FROM library/golang:1.12-alpine AS builder

RUN apk --no-cache add bash git make

WORKDIR /go/src/github.com/jwmarshall/check
COPY . .
RUN make build-prod

#
# App
FROM scratch

COPY --from=builder /go/src/github.com/jwmarshall/check/check .
CMD ["/check"]