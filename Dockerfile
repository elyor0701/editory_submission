FROM golang:1.20-bullseye as builder

RUN mkdir -p $GOPATH/src/gitlab.com/editory_submission/es_backend 
WORKDIR $GOPATH/src/gitlab.com/editory_submission/es_backend

COPY . ./
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/es_backend /

FROM alpine
COPY --from=builder es_backend .
ENTRYPOINT ["/es_backend"]
