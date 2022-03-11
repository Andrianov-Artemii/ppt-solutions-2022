FROM golang:latest
RUN mkdir checkit-bot
WORKDIR /checkit-bot
COPY . /checkit-bot
RUN go build . 
RUN chmod +x runner.sh
RUN echo $(ls -all)
ENTRYPOINT [ "/checkit-bot/runner.sh" ] 