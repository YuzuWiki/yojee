package client

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"io/ioutil"
	"net"
	"net/http"
)

// LookUpDNS by 1.1.1.1
type dNsBody struct {
	Status int `json:"status"`
	Answer []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		TTL  int    `json:"TTL"`
		Data string `json:"data"`
	} `json:"Answer"`
}

func LookUpDNS(host string) (err error, ip string) {
	if len(host) == 0 {
		return errors.New("invalid host"), ""
	}

	request, err := http.NewRequest(http.MethodGet, dnsQueryAPI+host, nil)
	if err != nil {
		return
	} else {
		request.Header.Set("Accept", "application/dns-json")
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		return errors.New(fmt.Sprintf("lookup dns fail, StatusCode is %d", response.StatusCode)), ""
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	body := dNsBody{}
	if err = json.Unmarshal(data, &body); err != nil {
		return
	}

	for idx := range body.Answer {
		answer := body.Answer[idx]
		if answer.Type == 1 {
			return nil, answer.Data
		}
	}
	return errors.New("not found IP"), ""
}

// ByPassSNI Transport.DialTLSContext
var cacheHosts = map[string]string{}

func resolveHost(host string) (string, error) {
	if v, isOk := cacheHosts[host]; isOk && v != "" {
		return v, nil
	}

	err, ip := LookUpDNS(host)
	if err != nil {
		return "", err
	} else {
		cacheHosts[host] = ip
	}

	return ip, nil
}

func DialTLSContext(ctx context.Context, network, addr string) (net.Conn, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	ip, err := resolveHost(host)
	if err != nil {
		return nil, err
	}

	return tls.Dial(
		network,
		net.JoinHostPort(ip, port),
		&tls.Config{
			InsecureSkipVerify: true,
			VerifyPeerCertificate: func(rawCerts [][]byte, _ [][]*x509.Certificate) error {

				roots := x509.NewCertPool()
				for _, rawCert := range rawCerts {
					c, err := x509.ParseCertificate(rawCert)
					if err != nil {
						return err
					}

					roots.AddCert(c)
				}

				cert, err := x509.ParseCertificate(rawCerts[0])
				if err != nil {
					return err
				}

				opts := x509.VerifyOptions{
					DNSName: cert.Subject.CommonName,
					Roots:   roots,
				}

				if _, err := cert.Verify(opts); err != nil {
					panic("Failed to verify certificate: " + err.Error())
					return err
				}
				return nil

			},
		})
}
