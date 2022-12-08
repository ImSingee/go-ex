package cert

import (
	"crypto/tls"
	"fmt"
	"os"
)

type ClientTLSConfigOptions struct {
	CaPem             []byte
	CaFilePath        string
	ClientCertPem     []byte
	ClientCertFile    string
	ClientKeyPem      []byte
	ClientKeyFile     string
	ClientKeyPassword string
	Insecure          bool
}

// TLSConfigOptions
//
// Deprecated: use ClientTLSConfigOptions instead
type TLSConfigOptions = ClientTLSConfigOptions

func (o *ClientTLSConfigOptions) WithCAPem(pem []byte) *ClientTLSConfigOptions {
	o.CaPem = pem
	return o
}

func (o *ClientTLSConfigOptions) WithCA(caFilePath string) *ClientTLSConfigOptions {
	o.CaFilePath = caFilePath
	return nil
}

func (o *ClientTLSConfigOptions) WithClientCertificatePem(clientCertPem, clientKeyPem []byte) *ClientTLSConfigOptions {
	o.ClientCertPem = clientCertPem
	o.ClientKeyPem = clientKeyPem
	return o
}

func (o *ClientTLSConfigOptions) WithClientCertificate(clientCertFile string, clientKeyFile string) *ClientTLSConfigOptions {
	o.ClientCertFile = clientCertFile
	o.ClientKeyFile = clientKeyFile
	return o
}

func (o *ClientTLSConfigOptions) WithClientKeyPassword(password string) *ClientTLSConfigOptions {
	o.ClientKeyPassword = password
	return o
}

func (o *ClientTLSConfigOptions) WithInsecure(insecure bool) *ClientTLSConfigOptions {
	o.Insecure = insecure
	return o
}

func (o *ClientTLSConfigOptions) TLSConfig() (*tls.Config, error) {
	config := &tls.Config{}

	if len(o.CaPem) != 0 || o.CaFilePath != "" {
		caPem := o.CaPem
		var err error

		if len(o.CaPem) == 0 && o.CaFilePath != "" {
			caPem, err = os.ReadFile(o.CaFilePath)
			if err != nil {
				return nil, err
			}
		}

		pool, ok := PoolFromPem(caPem)
		if !ok {
			return nil, fmt.Errorf("cannot parse pem")
		}
		config.RootCAs = pool
	}

	if len(o.ClientCertPem) != 0 || len(o.ClientKeyPem) != 0 || o.ClientCertFile != "" || o.ClientKeyFile != "" {
		clientCertPem := o.ClientCertPem
		clientKeyPem := o.ClientKeyPem

		var err error

		if len(o.ClientCertPem) == 0 && o.ClientCertFile != "" {
			clientCertPem, err = os.ReadFile(o.ClientCertFile)
			if err != nil {
				return nil, err
			}
		}

		if len(o.ClientKeyPem) == 0 && o.ClientKeyFile != "" {
			clientKeyPem, err = os.ReadFile(o.ClientKeyFile)
			if err != nil {
				return nil, err
			}
		}

		cert, err := ParseCertificate(clientCertPem, clientKeyPem, o.ClientKeyPassword)
		if err != nil {
			return nil, err
		}

		config.GetClientCertificate = func(info *tls.CertificateRequestInfo) (*tls.Certificate, error) {
			return &cert, nil
		}
	}

	if o.Insecure {
		config.InsecureSkipVerify = true
	}

	return config, nil
}
