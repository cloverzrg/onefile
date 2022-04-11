FROM alpine:latest
RUN apk update && apk add --no-cache ca-certificates tzdata \
    && rm -rf /var/cache/apk/*
COPY onefile /app/
ENV ENV prod
EXPOSE 80
RUN mkdir /app/data
VOLUME /app/data
WORKDIR /app
CMD ["/app/onefile"]