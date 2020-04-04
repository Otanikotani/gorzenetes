# Build
FROM golang:latest as build

ARG CI
ARG GIT_BRANCH
ARG SKIP_TEST

ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux

ADD . /app
WORKDIR /app

# run tests
RUN go test -timeout=30s  ./...
RUN go build -a -installsuffix cgo -o gorzenetes ./app && ls -la app/

# Run
# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=build /app/gorzenetes ./
COPY --from=build /app/art.txt ./

EXPOSE 8080

CMD ["/gorzenetes", "-bind", ":8080"]
