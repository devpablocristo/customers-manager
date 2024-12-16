package pkggrpcserver

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	defs "github.com/devpablocristo/customer-manager/pkg/grpc/server/defs"
)

// config estructura que implementa la interfaz defs.Config para el servidor
type config struct {
	host      string
	port      int
	tlsConfig *defs.TLSConfig
}

// newServerConfig crea una nueva configuración para el servidor gRPC
func newConfig(host string, port int, tlsConfig *defs.TLSConfig) defs.Config {
	return &config{
		host:      host,
		port:      port,
		tlsConfig: tlsConfig,
	}
}

func (c *config) GetHost() string {
	return c.host
}

func (c *config) SetHost(host string) {
	c.host = host
}

func (c *config) GetPort() int {
	return c.port
}

func (c *config) SetPort(port int) {
	c.port = port
}

func (c *config) GetTLSConfig() *defs.TLSConfig {
	return c.tlsConfig
}

func (c *config) SetTLSConfig(tlsConfig *defs.TLSConfig) {
	c.tlsConfig = tlsConfig
}

func (c *config) Validate() error {
	if c.port == 0 {
		return fmt.Errorf("gRPC server port is not configured")
	}
	return nil
}

// loadTLSConfig carga la configuración TLS para la conexión gRPC del servidor
func loadTLSConfig(tlsConfig *defs.TLSConfig) (*tls.Config, error) {
	certificate, err := tls.LoadX509KeyPair(tlsConfig.CertFile, tlsConfig.KeyFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(tlsConfig.CAFile)
	if err != nil {
		return nil, err
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, fmt.Errorf("failed to append CA certificates")
	}

	return &tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}, nil
}