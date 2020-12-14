FROM alpine:3.7
RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*
ADD linuxmicro /linuxmicro
WORKDIR /
ENTRYPOINT [ "/linuxmicro" ]
