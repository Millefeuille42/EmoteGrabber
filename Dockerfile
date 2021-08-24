FROM golang:1.16.0-alpine3.13

ADD ./grabber /grabber/
WORKDIR /grabber

RUN go install .

CMD grabber ./dev.env