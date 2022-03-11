FROM golang:latest
RUN mkdir checkit-bot
WORKDIR /checkit-bot
COPY . /checkit-bot

RUN go build .
ARG INPUT_KEYS=null
ENV RUN_KEYS=$INPUT_KEYS
CMD ppt-alert-bot $RUN_KEYS