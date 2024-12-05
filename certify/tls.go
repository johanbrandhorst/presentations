import (
	"crypto/x509"
	"io"
	"time"
)

type Config struct {
	...
	// GetCertificate returns a Certificate based on the given
	// ClientHelloInfo. It will only be called if the client supplies SNI
	// information or if Certificates is empty.
	GetCertificate func(*ClientHelloInfo) (*Certificate, error) // Go 1.4

	// GetClientCertificate, if not nil, is called when a server requests a
	// certificate from a client. If set, the contents of Certificates will
	// be ignored.
	GetClientCertificate func(*CertificateRequestInfo) (*Certificate, error) // Go 1.8
	...
}
