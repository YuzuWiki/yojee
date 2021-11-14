package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net"
	"net/http"
	"sync"
	"time"
)

type BypassSNIBlockingTransport struct {
	wrapped                http.RoundTripper
	antiSNIDetectTransport http.RoundTripper
	mu                     sync.Mutex
}

func (t *BypassSNIBlockingTransport) ensureWrappedTransport() http.RoundTripper {
	if t.wrapped == nil {
		return http.DefaultTransport
	}
	return t.wrapped
}

func (t *BypassSNIBlockingTransport) ensureAntiSNIDetectTransport() http.RoundTripper {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.antiSNIDetectTransport == nil {
		var v = new(http.Transport)
		v.DialTLSContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			host, port, err := net.SplitHostPort(addr)
			if err != nil {
				return nil, err
			}

			ip, err := resolveHostname(host)
			if err != nil {
				return nil, err
			}

			return tls.Dial(network, net.JoinHostPort(ip, port), &tls.Config{
				InsecureSkipVerify: true,
				VerifyPeerCertificate: func(certificate [][]byte, _ [][]*x509.Certificate) error {
					certs := make([]*x509.Certificate, len(certificate))
					for i, asn1Data := range certificate {
						cert, err := x509.ParseCertificate(asn1Data)
						if err != nil {
							return err
						}

						certs[i] = cert
					}

					opts := x509.VerifyOptions{
						DNSName:       host,
						Intermediates: x509.NewCertPool(),
					}

					for _, cert := range certs[1:] {
						opts.Intermediates.AddCert(cert)
					}

					cert := certs[0]
					_, err = cert.Verify(opts)
					if err != nil {
						return err
					}

					if time.Now().After(cert.NotAfter) {
						return errors.New("pixiv clitne: certification is expired")
					}

					if err = cert.VerifyHostname(host); err != nil {
						return err
					}
					return nil
				},
			})
		}
		t.antiSNIDetectTransport = v
	}
	return t.antiSNIDetectTransport
}

var (
	BlockedHostname = map[string]struct{}{
		"www.pixiv.net": {},
		"i.pximg.net":   {},
	}
)

func (t *BypassSNIBlockingTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	if _, isOk := BlockedHostname[req.URL.Host]; !isOk {
		return t.ensureWrappedTransport().RoundTrip(req)
	}
	return t.ensureAntiSNIDetectTransport().RoundTrip(req)
}

func (c *Client) BypassSNIBloccking() {
	c.Transport = &BypassSNIBlockingTransport{wrapped: c.Transport}
}
