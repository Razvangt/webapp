# Specifies a parent image
FROM golang:1.22.1-alpine
 
WORKDIR /app

COPY . .
RUN go mod download

COPY *.go ./

RUN go build -o /webapp

EXPOSE 8080

CMD [ "/webapp" ]
