package main

import "crypto/tls"

var tlsConfig = &tls.Config{
	ClientAuth: tls.RequireAndVerifyClientCert,
}

var h = VerifyClient(mux, "uw.com", "uw.co.uk")
