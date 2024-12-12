FROM node:23 as build-front

COPY front/package.json front/package-lock.json ./
RUN npm i
COPY front/ ./
RUN npm run build


FROM golang:1.23 as build-back

COPY back/ ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o /build/main main.go


FROM alpine:3.21

WORKDIR /go/bin/
COPY --from=build-front /build/ ./static/
COPY --from=build-back  /build/main ./main

ENV PORT=8080
EXPOSE 8080

HEALTHCHECK CMD wget --no-verbose --tries=1 --spider http://localhost:${PORT}/health || exit 1

ENTRYPOINT ["/go/bin/main", "--staticFilepath", "./static"]
