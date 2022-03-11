FROM golang:latest
RUN mkdir checkit-bot
WORKDIR /checkit-bot
COPY . /checkit-bot

RUN go build .
ARG INPUT_KEYS=null
ENV RUN_KEYS=$INPUT_KEYS
ENTRYPOINT ./ppt-alert-bot