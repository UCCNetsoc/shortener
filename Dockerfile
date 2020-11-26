FROM golang:1.15-alpine as DEV

WORKDIR /shortener

RUN GO111MODULE=on go get github.com/cortesi/modd/cmd/modd

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go install github.com/UCCNetsoc/shortener

CMD ["go", "run", "*.go"]

FROM alpine

WORKDIR /bin

COPY --from=DEV /go/bin/shortener ./shortener

CMD ["sh", "-c", "shortener -p"]