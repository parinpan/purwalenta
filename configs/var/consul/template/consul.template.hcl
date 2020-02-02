consul {
    address = "purwalenta_consul_server:8500"
    retry {
        enabled = true
        attempts = 20
        backoff = "3s"
    }
}

template {
	source      = "/etc/purwalenta/app.ctmpl"
	destination = "/etc/purwalenta/app.production.yaml"
}

template {
	source      = "/etc/purwalenta/app.ctmpl"
	destination = "/etc/purwalenta/app.staging.yaml"
}
