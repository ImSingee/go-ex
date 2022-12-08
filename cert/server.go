package cert

import (
	"crypto/tls"
	"fmt"
	"os"
)

type ServerTLSConfigOptions struct {
	ServerName        string
	ServerCertPem     []byte
	ServerCertFile    string
	ServerKeyPem      []byte
	ServerKeyFile     string
	ServerKeyPassword string
	ClientAuth        tls.ClientAuthType // for client auth
	CaPem             []byte             // for client auth
	CaFilePath        string             // for client auth
}

func (o *ServerTLSConfigOptions) WithServerName(serverName string) *ServerTLSConfigOptions {
	o.ServerName = serverName
	return o
}

func (o *ServerTLSConfigOptions) WithServerCertificatePem(serverCertPem, serverKeyPem []byte) *ServerTLSConfigOptions {
	o.ServerCertPem = serverCertPem
	o.ServerKeyPem = serverKeyPem
	return o
}

func (o *ServerTLSConfigOptions) WithServerCertificate(serverCertFile string, serverKeyFile string) *ServerTLSConfigOptions {
	o.ServerCertFile = serverCertFile
	o.ServerKeyFile = serverKeyFile
	return o
}

func (o *ServerTLSConfigOptions) WithServerKeyPassword(password string) *ServerTLSConfigOptions {
	o.ServerKeyPassword = password
	return o
}

func (o *ServerTLSConfigOptions) WithClientAuth(auth tls.ClientAuthType) *ServerTLSConfigOptions {
	o.ClientAuth = auth
	return o
}

func (o *ServerTLSConfigOptions) WithCAPem(pem []byte) *ServerTLSConfigOptions {
	o.CaPem = pem
	return o
}

func (o *ServerTLSConfigOptions) WithCA(caFilePath string) *ServerTLSConfigOptions {
	o.CaFilePath = caFilePath
	return nil
}

func (o *ServerTLSConfigOptions) TLSConfig() (*tls.Config, error) {
	config := &tls.Config{}

	config.ServerName = o.ServerName

	if len(o.ServerCertPem) != 0 || o.ServerCertFile != "" {
		serverCertPem := o.ServerCertPem
		var err error

		if len(o.ServerCertPem) == 0 && o.ServerCertFile != "" {
			serverCertPem, err = os.ReadFile(o.ServerCertFile)
			if err != nil {
				return nil, fmt.Errorf("read server cert file: %w", err)
			}
		}

		serverKeyPem := o.ServerKeyPem
		if len(o.ServerKeyPem) == 0 && o.ServerKeyFile != "" {
			serverKeyPem, err = os.ReadFile(o.ServerKeyFile)
			if err != nil {
				return nil, fmt.Errorf("read server key file: %w", err)
			}
		}

		cert, err := tls.X509KeyPair(serverCertPem, serverKeyPem)
		if err != nil {
			return nil, fmt.Errorf("load server cert: %w", err)
		}

		config.Certificates = []tls.Certificate{cert}
	}

	// --- below setup client auth

	config.ClientAuth = o.ClientAuth

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
		config.ClientCAs = pool
	}
	return config, nil
}
