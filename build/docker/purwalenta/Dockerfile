FROM alpine:latest
LABEL maintainer="Fachrin Aulia <fachrinfan@gmail.com>"

RUN mkdir -p /var/consul/template
RUN mkdir -p /etc/purwalenta

RUN wget https://releases.hashicorp.com/consul-template/0.24.1/consul-template_0.24.1_linux_amd64.zip
RUN unzip consul-template_0.24.1_linux_amd64.zip; rm consul-template_0.24.1_linux_amd64.zip
RUN mv consul-template /usr/bin/consul-template

COPY ./configs/etc/purwalenta/app.ctmpl /etc/purwalenta/app.ctmpl
COPY ./configs/etc/purwalenta/app.yaml /etc/purwalenta/app.yaml

COPY ./configs/var/consul/template/consul.template.hcl /var/consul/template/consul.template.hcl
COPY ./cmd/purwalenta/purwalenta-bin /usr/bin/purwalenta
COPY ./build/docker/purwalenta/purwalenta-app.sh /usr/bin/purwalenta-app

ENV PATH="/usr/bin:${PATH}"
