FROM golang:latest
RUN mkdir checkit-bot
WORKDIR /checkit-bot
COPY . /checkit-bot
ARG INPUT_KEYS=null
ENV RUN_KEYS=$INPUT_KEYS
RUN echo $INPUT_KEYS
RUN go build .
CMD ./ppt-alert-bot $RUN_KEYS