FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

ADD consignment.json /app/consignment.json
ADD shippy-consignment-cli /app/shippy-consignment-cli

CMD ["./shippy-consignment-cli"]