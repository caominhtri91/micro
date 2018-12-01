FROM golang:latest

EXPOSE 6767
ADD accountservice-linux-amd64 /

RUN chmod +x /accountservice-linux-amd64
CMD ["/accountservice-linux-amd64"]
