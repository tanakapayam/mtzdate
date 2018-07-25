# building...
FROM golang:alpine as builder

ENV APP="mtzdate" \
    CGO_ENABLED=0 \
    TERM="xterm"

WORKDIR /go/src/github.com/tanakapayam/"$APP"
COPY . .

RUN apk --update --no-cache add \
        bash \
        git \
        make \
        ncurses \
        tzdata \
    && make install

# minimal image
FROM alpine:3.8

ENV APP="mtzdate" \
    MTZDATE_LOOP="" \
    MTZDATE_FLAGS="" \
    MTZDATE_TIMEZONES="" \
    MTZDATE_WORKDAYS="Mon,Tue,Wed,Thu,Fri" \
    MTZDATE_GREEN_HOURS="8-17" \
    MTZDATE_YELLOW_HOURS="7-8,17-18" \
    MTZDATE_FAINT_HOURS="0-7,22-24" \
    MTZDATE_FORMAT="dfc" \
    TERM="xterm"

RUN apk --update --no-cache add \
        ncurses \
    && addgroup -S "$APP" \
    && adduser -D -S -G "$APP" -H -h "/$APP" "$APP" \
    && rm -rf /var/cache/apk/*

COPY --from=builder \
    /go/bin/"$APP" \
    /usr/bin/"$APP"
COPY --from=builder \
    /usr/share/zoneinfo \
    /usr/share/zoneinfo

USER "$APP"
WORKDIR "/$APP"
ENTRYPOINT ["mtzdate"]

## trapped by mtzdate
STOPSIGNAL "SIGINT"
