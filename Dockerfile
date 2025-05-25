FROM golang:1.24

WORKDIR /app

EXPOSE 8080
 CMD [ "/main" ]