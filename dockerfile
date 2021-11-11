FROM amd64/golang:1.17
WORKDIR /go/src/github.com/mgould1799/ip-checker
COPY . .
RUN go mod download 
RUN go build
EXPOSE 8080
CMD ["./ip-checker"]