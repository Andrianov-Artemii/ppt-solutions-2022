FROM golang:latest
RUN mkdir checkit-bot
WORKDIR /checkit-bot
COPY . /checkit-bot

RUN go build .
ENV RUN_KEYS=null
CMD ./ppt-alert-bot $RUN_KEYS