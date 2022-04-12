FROM --platform=$BUILDPLATFORM golang:alpine as builder
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH
ARG DRONE_TAG
ENV CGO_ENABLED 0
ENV GOOS $TARGETOS
ENV GOARCH $TARGETARCH
RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM, GOOS $GOOS, GOARCH $GOARCH"
RUN apk update && apk add --no-cache git build-base
COPY . /go/src/onefile
WORKDIR /go/src/onefile
RUN go build -o onefile

FROM --platform=$TARGETPLATFORM alpine:latest
RUN apk update && apk add --no-cache ca-certificates tzdata \
    && rm -rf /var/cache/apk/*
COPY --from=builder /go/src/onefile/onefile /app/onefile
ENV ENV prod
EXPOSE 80
VOLUME /app/data
WORKDIR /app
CMD ["/app/onefile"]