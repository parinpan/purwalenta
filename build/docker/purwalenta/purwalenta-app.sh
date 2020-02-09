#!/bin/sh

consul-template -config /var/consul/template/consul.template.hcl -once &
purwalenta
