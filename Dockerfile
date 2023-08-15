##
## Build stage
##
FROM golang:1.21.0-alpine3.18 AS builder
RUN apk update && apk upgrade
WORKDIR /src
COPY . .

#RUN go mod download
RUN GOOS=linux go build -o /tfs-telbot .

##
## Final image stage
##
FROM alpine:3.18.3
WORKDIR /
COPY --from=builder /tfs-telbot /tfs-telbot

ENTRYPOINT ["/tfs-telbot"]