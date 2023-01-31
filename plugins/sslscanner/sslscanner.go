package sslscanner

import "malgo"

import (
	"crypto/tls"
	"fmt"
)

// SSLScannerPlugin ssl sertifikasının geçerliliğini kontrol eder
type SSLScannerPlugin struct {
	// SSLScannerPlugin için özel alanlar
}

func (p *SSLScannerPlugin) GetName() string {
	return "SSL Scanner"
}

func (p *SSLScannerPlugin) Symbol() string {
	return "SSLScanner"
}

func (p *SSLScannerPlugin) Register() main.MyPlugin {
	fmt.Println("sslscanner plugin registered")
	return p
}

// Run sertifikasın geçerliliğini kontrol eder
func (p *SSLScannerPlugin) Run(host string) {
	fmt.Println("sslscanner plugin running")
	conn, err := tls.Dial("tcp", host, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		fmt.Println(cert.Subject)
	}
}
