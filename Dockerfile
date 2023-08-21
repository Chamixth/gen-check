FROM golang:1.19

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go get

COPY *.go ./

RUN go build -o /chaithApp

EXPOSE 8088

CMD [ "/chaithApp" ]