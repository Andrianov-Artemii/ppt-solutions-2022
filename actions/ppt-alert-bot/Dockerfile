FROM golang:latest
RUN mkdir checkit-bot
WORKDIR /checkit-bot
COPY . /checkit-bot
RUN go build . ; chmod +x runner.sh
ENTRYPOINT [ "./runner.sh" ] 