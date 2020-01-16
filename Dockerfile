# Stage 1
FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
RUN mkdir /build 
ADD . /build/
WORKDIR /build
#Get all dependices within the GoAPplication
RUN go get -d -v
RUN go build -o cicdexample .
# Stage 2
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/ /app/
WORKDIR /app
EXPOSE 8081
CMD ["./cicdexample"]