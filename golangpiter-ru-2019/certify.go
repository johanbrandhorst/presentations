package main

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/johanbrandhorst/certify"
	"github.com/johanbrandhorst/certify/issuers/vault"
)

// Issuer is the interface that must be implemented
// by certificate issuers.
type Issuer interface {
	Issue(ctx context.Context, cn string, conf *certify.CertConfig) (*tls.Certificate, error)
}

func main() {
	issuer := &vault.Issuer{
		// URL, Token and Role are required
		URL: &url.URL{
			Scheme: "https",
			Host:   "my-local-vault-instance.com",
		},
		Token: "myVaultToken",
		Role:  "myVaultRole",
		// TimeToLive configures how long new certificates should
		// be valid for. Defaults to the max of the role.
		TimeToLive: time.Hour * 24 * 30,
		// OtherSubjectAlternativeNames defines custom OID/UTF8-string SANs.
		OtherSubjectAlternativeNames: nil,
	}

	c := &certify.Certify{
		CommonName: "uw.co.uk",
		Issuer:     issuer,
		// Cache certificates in memory.
		Cache: certify.NewMemCache(),
		// Refresh cached certificates when < 24H left before expiry.
		RenewBefore: 24 * time.Hour,
		// Optionally configure properties of the certificates.
		CertConfig: &certify.CertConfig{
			SubjectAlternativeNames:   []string{"uw.com"},
			IPSubjectAlternativeNames: []net.IP{net.IPv6loopback},
			// Can also specify a custom private key generator, defaults to ECDSA p256.
		},
	}

	s := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate:       c.GetCertificate,
			GetClientCertificate: c.GetClientCertificate,
		},
	}
	s.ListenAndServeTLS("", "")
}
