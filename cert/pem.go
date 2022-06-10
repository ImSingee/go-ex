package cert

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func ParseCertificateFile(certFile, keyFile string, password ...string) (tls.Certificate, error) {
	certPEM, err := os.ReadFile(certFile)
	if err != nil {
		return tls.Certificate{}, err
	}
	keyPEM, err := os.ReadFile(keyFile)
	if err != nil {
		return tls.Certificate{}, err
	}

	return ParseCertificate(certPEM, keyPEM, password...)
}

// ParseCertificate 解析 PEM 证书文件
func ParseCertificate(certPEM, keyPEM []byte, password ...string) (tls.Certificate, error) {
	if len(password) == 0 || password[0] == "" {
		return tls.X509KeyPair(certPEM, keyPEM)
	}

	var v *pem.Block
	var pkey []byte

	for {
		v, keyPEM = pem.Decode(keyPEM)
		if v == nil {
			break
		}
		if v.Type == "RSA PRIVATE KEY" {
			if x509.IsEncryptedPEMBlock(v) {
				pkey, _ = x509.DecryptPEMBlock(v, []byte(password[0]))
				pkey = pem.EncodeToMemory(&pem.Block{
					Type:  v.Type,
					Bytes: pkey,
				})
			} else {
				pkey = pem.EncodeToMemory(v)
			}

			break
		}
	}

	return tls.X509KeyPair(certPEM, pkey)
}

func PoolFromPem(pem []byte) (pool *x509.CertPool, ok bool) {
	pool = x509.NewCertPool()
	ok = pool.AppendCertsFromPEM(pem)

	return pool, ok
}

func PoolFromPemFile(pemFile string) (*x509.CertPool, error) {
	pemData, err := os.ReadFile(pemFile)
	if err != nil {
		return nil, err
	}

	pool, ok := PoolFromPem(pemData)
	if !ok {
		return nil, fmt.Errorf("cannot parse pem")
	}

	return pool, nil
}
