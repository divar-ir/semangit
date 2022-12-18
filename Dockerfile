FROM alpine:3.17.0
ARG HTTPS_PROXY
RUN apk update && apk add git bash libc6-compat
COPY semangit /usr/bin/
RUN chmod +x /usr/bin/semangit
