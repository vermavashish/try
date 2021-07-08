FROM golang:onbuild
 
RUN mkdir -p /try
 
WORKDIR /try

RUN go mod init github.com/vermavashish/try
 
ADD . /try
 
RUN go build ./main.go
 
CMD ["./main"]