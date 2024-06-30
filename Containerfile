FROM docker.io/alpine:3.20

LABEL org.opencontainers.image.source https://github.com/Marco98/ytpodproxy
ENTRYPOINT ["/usr/local/bin/ytpodproxy"]

RUN apk add --no-cache \
    yt-dlp \
    ffmpeg

COPY ytpodproxy /usr/local/bin/ytpodproxy
