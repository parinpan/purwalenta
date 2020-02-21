#!/bin/sh

purwalenta >> /var/log/purwalenta/upstart.log 2>&1
consul-template -config /var/consul/template/consul.template.hcl >> /var/log/purwalenta/consul.log 2>&1
