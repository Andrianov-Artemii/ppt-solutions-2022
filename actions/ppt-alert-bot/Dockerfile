FROM golang:latest
RUN mkdir checkit-bot
WORKDIR /checkit-bot
COPY . /checkit-bot
ENV RUN_KEYS=null
RUN echo $RUN_KEYS
RUN go build .
CMD ./ppt-alert-bot $RUN_KEYS