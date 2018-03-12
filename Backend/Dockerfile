FROM golang
ADD . go/src/github.com/gorilla/mux
WORKDIR go/src/github.com/gorilla/mux
RUN go get go/src/github.com/gorilla/mux
RUN go install
RUN go build -o main
ENV PORT 8080
EXPOSE 8080
CMD [“/main”]