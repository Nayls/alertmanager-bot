FROM golang:1.14-alpine3.12 AS compile-image

ARG CI_COMMIT_REF_SLUG
ARG CI_PIPELINE_URL
ARG CI_COMMIT_TITLE
ARG CI_COMMIT_SHA
ARG GIT_USER_NAME
ARG GIT_USER_EMAIL

WORKDIR /go/src/github.com/Nayls/alertmanager-bot
COPY . .
RUN CGO_ENABLED=0 GOOS=linux \
    go build \
    -mod vendor \
    -a \
    -ldflags "-X 'main.BuildDate=$(date)' \
            -X 'main.CI_COMMIT_REF_SLUG=${CI_COMMIT_REF_SLUG}' \
            -X 'main.CI_PIPELINE_URL=${CI_PIPELINE_URL}'  \
            -X 'main.CI_COMMIT_TITLE=${CI_COMMIT_TITLE}' \
            -X 'main.CI_COMMIT_SHA=${CI_COMMIT_SHA}' \
            -X 'main.GIT_USER_NAME=${GIT_USER_NAME}' \
            -X 'main.GIT_USER_EMAIL=${GIT_USER_EMAIL}'" \
    -installsuffix cgo -o ./bin/alertmanager-bot ./main.go

FROM alpine:3.12 AS runtime-image

RUN apk add --no-cache --update \
        ca-certificates \
        tzdata \
    && rm -rf /var/cache/apk/*

WORKDIR /alertmanager-bot
COPY --from=compile-image /go/src/github.com/Nayls/alertmanager-bot/config.yaml ./
COPY --from=compile-image /go/src/github.com/Nayls/alertmanager-bot/bin/alertmanager-bot /usr/local/bin
RUN chmod +x /usr/local/bin/alertmanager-bot

VOLUME /alertmanager-bot

EXPOSE 8080
CMD ["alertmanager-bot"]