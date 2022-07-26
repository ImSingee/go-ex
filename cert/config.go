package cert

import (
	"crypto/tls"
	"fmt"
	"os"
)

type TLSConfigOptions struct {
	CaPem             []byte
	CaFilePath        string
	ClientCertPem     []byte
	ClientCertFile    string
	ClientKeyPem      []byte
	ClientKeyFile     string
	ClientKeyPassword string
	Insecure          bool
}

func (o *TLSConfigOptions) WithCAPem(pem []byte) *TLSConfigOptions {
	o.CaPem = pem
	return o
}

func (o *TLSConfigOptions) WithCA(caFilePath string) *TLSConfigOptions {
	o.CaFilePath = caFilePath
	return nil
}

func (o *TLSConfigOptions) WithClientCertificatePem(clientCertPem, clientKeyPem []byte) *TLSConfigOptions {
	o.ClientCertPem = clientCertPem
	o.ClientKeyPem = clientKeyPem
	return o
}

func (o *TLSConfigOptions) WithClientCertificate(clientCertFile string, clientKeyFile string) *TLSConfigOptions {
	o.ClientCertFile = clientCertFile
	o.ClientKeyFile = clientKeyFile
	return o
}

func (o *TLSConfigOptions) WithClientKeyPassword(password string) *TLSConfigOptions {
	o.ClientKeyPassword = password
	return o
}

func (o *TLSConfigOptions) WithInsecure(insecure bool) *TLSConfigOptions {
	o.Insecure = insecure
	return o
}

func (o *TLSConfigOptions) TLSConfig() (*tls.Config, error) {
	config := &tls.Config{}

	if len(o.CaPem) != 0 || o.CaFilePath != "" {
		var caPem []byte
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
		var clientCertPem, clientKeyPem []byte
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
