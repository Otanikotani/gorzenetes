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
RUN go build -a -installsuffix cgo -o gorzenetes ./app && ls -la

# Run
FROM scratch

COPY --from=build /app/gorzenetes ./

EXPOSE 80

CMD ["/gorzenetes", "-bind", ":80"]
