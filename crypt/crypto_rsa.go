// Package crypt provides cryptographic utilities including RSA encryption,
// AES encryption, hashing functions, and certificate generation.
package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"os"
	"time"
)

// KeyFormat represents the format of RSA private keys.
type KeyFormat int

const (
	// PKCS1 represents the PKCS#1 format for RSA private keys.
	PKCS1 KeyFormat = iota

	// PKCS8 represents the PKCS#8 format for private keys.
	PKCS8
)

// CertificateConfig holds configuration information for generating X.509 certificates.
type CertificateConfig struct {
	// Country is the country name (C) field in the certificate subject.
	Country string

	// Province is the state or province name (ST) field in the certificate subject.
	Province string

	// Locality is the locality name (L) field in the certificate subject.
	Locality string

	// Organization is the organization name (O) field in the certificate subject.
	Organization string

	// OrganizationalUnit is the organizational unit name (OU) field in the certificate subject.
	OrganizationalUnit string

	// CommonName is the common name (CN) field in the certificate subject.
	CommonName string

	// NotBefore is the time when the certificate becomes valid.
	NotBefore time.Time

	// NotAfter is the time when the certificate expires.
	NotAfter time.Time

	// SerialNumber is the serial number for the certificate.
	// If nil, a random serial number will be generated.
	SerialNumber *big.Int
}

func DefaultCertificateConfig() *CertificateConfig {
	now := time.Now()
	return &CertificateConfig{
		Country:            "CN",
		Province:           "Beijing",
		Locality:           "Beijing",
		Organization:       "Minzhan",
		OrganizationalUnit: "IT Department",
		CommonName:         "minzhan.com",
		NotBefore:          now,
		NotAfter:           now.AddDate(1, 0, 0),
	}
}

// GenerateKey generates an RSA private key with the specified bit size.
// It returns an error if the key generation fails.
func GenerateKey(bits int) (*rsa.PrivateKey, error) {
	if bits < 512 {
		return nil, errors.New("key size must be at least 512 bits")
	}

	return rsa.GenerateKey(rand.Reader, bits)
}

// SavePrivateKey saves an RSA private key to a file in the specified format.
// It supports both PKCS#1 and PKCS#8 formats.
// Returns an error if the file cannot be created or the key cannot be encoded.
func SavePrivateKey(filePath string, privateKey *rsa.PrivateKey, format KeyFormat) error {
	if privateKey == nil {
		return errors.New("private key cannot be nil")
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var blockType string
	var bytes []byte

	switch format {
	case PKCS1:
		blockType = "RSA PRIVATE KEY"
		bytes = x509.MarshalPKCS1PrivateKey(privateKey)
	case PKCS8:
		blockType = "PRIVATE KEY"
		bytes, err = x509.MarshalPKCS8PrivateKey(privateKey)
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported key format")
	}

	return pem.Encode(file, &pem.Block{Type: blockType, Bytes: bytes})
}

// LoadPrivateKey loads an RSA private key from a PEM-encoded file.
// It supports both PKCS#1 and PKCS#8 formats.
// Returns the private key and an error if loading fails.
func LoadPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	var key *rsa.PrivateKey

	switch block.Type {
	case "RSA PRIVATE KEY":
		// PKCS#1 format
		key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	case "PRIVATE KEY":
		// PKCS#8 format
		privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}

		var ok bool
		key, ok = privateKey.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("not an RSA private key")
		}
	default:
		return nil, errors.New("unsupported private key format")
	}

	if err != nil {
		return nil, err
	}

	return key, nil
}

// GenerateCertificate generates an X.509 certificate and saves it to a file.
// If config is nil, a default configuration will be used.
// Returns an error if the certificate generation or file operations fail.
func GenerateCertificate(filePath string, key *rsa.PrivateKey, config *CertificateConfig) error {
	if key == nil {
		return errors.New("private key cannot be nil")
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if config == nil {
		config = DefaultCertificateConfig()
	}

	// Generate a random serial number if not provided
	if config.SerialNumber == nil {
		config.SerialNumber, err = rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
		if err != nil {
			return err
		}
	}

	// Build certificate template
	template := &x509.Certificate{
		SerialNumber:          config.SerialNumber,
		NotBefore:             config.NotBefore,
		NotAfter:              config.NotAfter,
		IsCA:                  false,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		BasicConstraintsValid: true,
		Subject: pkix.Name{
			Country:            []string{config.Country},
			Province:           []string{config.Province},
			Locality:           []string{config.Locality},
			Organization:       []string{config.Organization},
			OrganizationalUnit: []string{config.OrganizationalUnit},
			CommonName:         config.CommonName,
		},
	}

	// Generate and encode certificate
	certBytes, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	if err != nil {
		return err
	}

	return pem.Encode(file, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
}

// LoadCertificate loads an X.509 certificate from a PEM-encoded file.
// Returns the certificate and an error if loading fails.
func LoadCertificate(filePath string) (*x509.Certificate, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	if block.Type != "CERTIFICATE" {
		return nil, errors.New("not a certificate")
	}

	return x509.ParseCertificate(block.Bytes)
}
