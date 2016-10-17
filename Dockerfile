FROM buildpack-deps:jessie

RUN mkdir -p /opt/microservices
COPY . /opt/microservices/
EXPOSE 8080
WORKDIR /opt/microservices

RUN wget -qO- https://github.com/amalgam8/amalgam8/releases/download/v0.4.0/a8sidecar.sh | sh

ENTRYPOINT ["a8sidecar", "--supervise", "./stuff"]
