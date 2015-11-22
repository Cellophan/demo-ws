FROM golang:alpine

COPY main.go /app/
RUN  cd /app &&\
     go build -v
CMD /app/app
