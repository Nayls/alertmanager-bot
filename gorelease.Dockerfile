FROM alpine:3.12 AS gorelease-image

RUN apk add --no-cache --update \
        ca-certificates \
        tzdata \
    && rm -rf /var/cache/apk/*

WORKDIR /alertmanager-bot
COPY ./ ./
COPY ./alertmanager-bot /usr/local/bin
RUN chmod +x /usr/local/bin/alertmanager-bot

VOLUME /alertmanager-bot

EXPOSE 8080
CMD ["alertmanager-bot"]