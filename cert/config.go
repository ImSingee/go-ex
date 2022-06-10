package cert

import "crypto/tls"

type TLSConfigOptions struct {
	CaFilePath        string
	ClientCertFile    string
	ClientKeyFile     string
	ClientKeyPassword string
	Insecure          bool
}

func (o *TLSConfigOptions) WithCA(caFilePath string) *TLSConfigOptions {
	o.CaFilePath = caFilePath
	return nil
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

	if o.CaFilePath != "" {
		pool, err := PoolFromPemFile(o.CaFilePath)
		if err != nil {
			return nil, err
		}
		config.RootCAs = pool
	}

	if o.ClientCertFile != "" {
		cert, err := ParseCertificateFile(o.ClientCertFile, o.ClientKeyFile, o.ClientKeyPassword)
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
