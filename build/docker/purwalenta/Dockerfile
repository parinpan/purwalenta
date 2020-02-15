FROM golang:alpine AS build-env

LABEL maintainer="Fachrin Aulia <fachrinfan@gmail.com>"
ENV INSTALL_DIRECTORY="/usr/bin"
ENV PATH="/usr/bin:${PATH}"

RUN rm -rf /var/cache/apk/* && rm -rf /tmp/*
RUN apk --no-cache add build-base git bzr mercurial gcc
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/purwalenta/purwalenta
COPY . .

#RUN dep ensure -v
RUN go build -o purwalenta-bin ./cmd/purwalenta/.

######################################################################################################################
FROM alpine:latest

LABEL maintainer="Fachrin Aulia <fachrinfan@gmail.com>"
ENV TZ="Asia/Jakarta"
ENV PATH="/usr/bin:${PATH}"

RUN rm -rf /var/cache/apk/* && rm -rf /tmp/*
RUN apk update && apk upgrade && apk add ca-certificates && apk add --update tzdata && apk add ssmtp

RUN mkdir -p /var/consul/template
RUN mkdir -p /etc/purwalenta

RUN wget https://releases.hashicorp.com/consul-template/0.24.1/consul-template_0.24.1_linux_amd64.zip
RUN unzip consul-template_0.24.1_linux_amd64.zip; rm consul-template_0.24.1_linux_amd64.zip
RUN mv consul-template /usr/bin/consul-template

WORKDIR /app
COPY . .

COPY ./configs/etc/purwalenta/app.ctmpl /etc/purwalenta/app.ctmpl
COPY ./configs/etc/purwalenta/app.yaml /etc/purwalenta/app.yaml
COPY ./configs/var/consul/template/consul.template.hcl /var/consul/template/consul.template.hcl
COPY ./configs/var/email/template/. /var/email/template/.
COPY ./build/docker/purwalenta/purwalenta-app.sh /usr/bin/purwalenta-app
COPY --from=build-env /go/src/github.com/purwalenta/purwalenta/purwalenta-bin /usr/bin/purwalenta
