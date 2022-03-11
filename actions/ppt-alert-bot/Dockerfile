FROM golang:latest
RUN mkdir checkit-bot
WORKDIR /checkit-bot
COPY . /checkit-bot

RUN go build .
CMD ./ppt-alert-bot $1