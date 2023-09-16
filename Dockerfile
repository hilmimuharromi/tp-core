FROM node:alpine as build-node
RUN apk --no-cache --virtual build-dependencies add

WORKDIR /workdir
COPY web/ .
RUN yarn install
RUN yarn build

FROM golang:alpine as build-go

ENV GOPATH ""

RUN apk update && apk upgrade

ADD go.mod go.sum ./
ADD .env ./
RUN go mod download
ADD . .
COPY --from=build-node /workdir/build ./web/build
RUN go build -o main .

# FROM alpine
# COPY --from=build-go /main .

EXPOSE 8080

CMD ["./main", "run"]