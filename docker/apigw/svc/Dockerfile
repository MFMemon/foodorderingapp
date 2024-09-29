FROM ubuntu:jammy
RUN apt-get update && apt-get install -y curl iputils-ping
WORKDIR /app
COPY ./cmd/apigateway/apigw .
RUN chmod +x apigw
# ENTRYPOINT [ "tail", "-f", "/dev/null" ]
ENTRYPOINT [ "./apigw" ]

