FROM alpine:latest
WORKDIR /
COPY main .
RUN apk add ca-certificates
CMD [ "./main" ]