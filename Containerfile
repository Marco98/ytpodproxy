FROM scratch

LABEL org.opencontainers.image.source https://github.com/Marco98/ytpodproxy
ENTRYPOINT ["/ytpodproxy"]

COPY ytpodproxy /ytpodproxy
