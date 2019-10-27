package main

import "crypto/tls"

var tlsConfig = &tls.Config{
	ClientAuth: tls.RequireAndVerifyClientCert,
}

var h = VerifyClient(mux, "mystore.com", "app.mystore.com")
