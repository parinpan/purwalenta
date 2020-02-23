#!/bin/sh

consul-template -config /var/consul/template/consul.template.hcl >> /var/log/purwalenta/consul.log 2>&1 &
purwalenta >> /var/log/purwalenta/upstart.log 2>&1
