FROM golang:latest
RUN mkdir checkit-bot
WORKDIR /checkit-bot
COPY . /checkit-bot
RUN go build . 
RUN chmod +x runner.sh
RUN echo $(ls -all)
RUN echo $(pwd)
ENTRYPOINT [ "/checkit-bot/runner.sh" ] 