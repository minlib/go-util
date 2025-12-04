package crypt

import (
	"math/big"
	"os"
	"testing"
	"time"
)

func TestGenerateKey(t *testing.T) {
	tests := []struct {
		name    string
		bits    int
		wantErr bool
	}{
		{
			name:    "Valid 2048 bits key",
			bits:    2048,
			wantErr: false,
		},
		{
			name:    "Valid 1024 bits key",
			bits:    1024,
			wantErr: false,
		},
		{
			name:    "Invalid key size",
			bits:    256,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateKey(tt.bits)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got.N.BitLen() != tt.bits {
				t.Errorf("GenerateKey() = %v bits, want %v bits", got.N.BitLen(), tt.bits)
			}
		})
	}
}

func TestGeneratePrivateKeyAndCertificate(t *testing.T) {
	privatePath := "./private.key"
	certPath := "./cert.pem"

	defer os.Remove(privatePath)
	defer os.Remove(certPath)

	key, err := GenerateKey(2048)
	if err != nil {
		t.Fatalf("Failed to generate test key: %v", err)
	}

	if err = SavePrivateKey(privatePath, key, PKCS8); err != nil {
		t.Errorf("Failed to save private key: %v", err)
	}

	if err := GenerateCertificate(certPath, key, nil); err != nil {
		t.Errorf("Failed to generate certificate: %v", err)
	}
}

func TestSavePrivateKeyAndLoadPrivateKey(t *testing.T) {
	// Test file paths
	pkcs1File := "/tmp/test_private_key_pkcs1.key"
	pkcs8File := "/tmp/test_private_key_pkcs8.key"

	// Clean up test files
	defer os.Remove(pkcs1File)
	defer os.Remove(pkcs8File)

	// Generate a test key
	key, err := GenerateKey(2048)
	if err != nil {
		t.Fatalf("Failed to generate test key: %v", err)
	}

	tests := []struct {
		name     string
		filePath string
		format   KeyFormat
		wantErr  bool
	}{
		{
			name:     "Save PKCS1 format",
			filePath: pkcs1File,
			format:   PKCS1,
			wantErr:  false,
		},
		{
			name:     "Save PKCS8 format",
			filePath: pkcs8File,
			format:   PKCS8,
			wantErr:  false,
		},
		{
			name:     "Invalid format",
			filePath: "/tmp/test_invalid.pem",
			format:   KeyFormat(99), // Invalid format
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test saving
			err := SavePrivateKey(tt.filePath, key, tt.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("SavePrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Skip loading test for invalid format case
			if tt.wantErr {
				return
			}

			// Test loading
			loadedKey, err := LoadPrivateKey(tt.filePath)
			if err != nil {
				t.Errorf("LoadPrivateKey() error = %v", err)
				return
			}

			// Verify the loaded key matches the original
			if loadedKey.N.Cmp(key.N) != 0 {
				t.Error("Loaded key does not match original key")
			}
		})
	}

	// Test with nil key
	t.Run("Save nil key", func(t *testing.T) {
		err := SavePrivateKey("/tmp/test_nil.key", nil, PKCS1)
		if err == nil {
			t.Error("SavePrivateKey() with nil key should return error")
		}
	})
}

func TestGenerateCertificateAndLoadCertificate(t *testing.T) {
	// Generate a test key
	key, err := GenerateKey(2048)
	if err != nil {
		t.Fatalf("Failed to generate test key: %v", err)
	}

	// Test file path
	certFile := "/tmp/test_certificate.pem"
	defer os.Remove(certFile)

	// Test with default config
	t.Run("Generate certificate with default config", func(t *testing.T) {
		err := GenerateCertificate(certFile, key, nil)
		if err != nil {
			t.Errorf("GenerateCertificate() with default config error = %v", err)
			return
		}

		// Test loading
		cert, err := LoadCertificate(certFile)
		if err != nil {
			t.Errorf("LoadCertificate() error = %v", err)
			return
		}

		// Verify certificate properties
		if cert.Subject.CommonName != "minzhan.com" {
			t.Errorf("Certificate CommonName = %v, want minzhan.com", cert.Subject.CommonName)
		}
	})

	// Test with custom config
	t.Run("Generate certificate with custom config", func(t *testing.T) {
		now := time.Now()
		config := &CertificateConfig{
			Country:            "US",
			Province:           "California",
			Locality:           "San Francisco",
			Organization:       "Test Org",
			OrganizationalUnit: "Test Unit",
			CommonName:         "test.example.com",
			NotBefore:          now,
			NotAfter:           now.AddDate(0, 0, 365), // 1 year
			SerialNumber:       big.NewInt(987654321),
		}

		customCertFile := "/tmp/test_certificate_custom.pem"
		defer os.Remove(customCertFile)

		err := GenerateCertificate(customCertFile, key, config)
		if err != nil {
			t.Errorf("GenerateCertificate() with custom config error = %v", err)
			return
		}

		// Test loading
		cert, err := LoadCertificate(customCertFile)
		if err != nil {
			t.Errorf("LoadCertificate() error = %v", err)
			return
		}

		// Verify certificate properties
		if cert.Subject.CommonName != "test.example.com" {
			t.Errorf("Certificate CommonName = %v, want test.example.com", cert.Subject.CommonName)
		}
		if cert.Subject.Country[0] != "US" {
			t.Errorf("Certificate Country = %v, want US", cert.Subject.Country[0])
		}
		if cert.SerialNumber.Cmp(big.NewInt(987654321)) != 0 {
			t.Errorf("Certificate SerialNumber = %v, want 987654321", cert.SerialNumber)
		}
	})

	// Test with nil key
	t.Run("Generate certificate with nil key", func(t *testing.T) {
		err := GenerateCertificate("/tmp/test_cert_nil.pem", nil, nil)
		if err == nil {
			t.Error("GenerateCertificate() with nil key should return error")
		}
	})
}

func TestLoadPrivateKeyErrors(t *testing.T) {
	// Test non-existent file
	_, err := LoadPrivateKey("/tmp/non_existent_file.pem")
	if err == nil {
		t.Error("LoadPrivateKey() with non-existent file should return error")
	}

	// Create an invalid PEM file
	invalidFile := "/tmp/invalid.pem"
	defer os.Remove(invalidFile)

	err = os.WriteFile(invalidFile, []byte("invalid pem content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create invalid PEM file: %v", err)
	}

	_, err = LoadPrivateKey(invalidFile)
	if err == nil {
		t.Error("LoadPrivateKey() with invalid PEM should return error")
	}
}

func TestLoadCertificateErrors(t *testing.T) {
	// Test non-existent file
	_, err := LoadCertificate("/tmp/non_existent_cert.pem")
	if err == nil {
		t.Error("LoadCertificate() with non-existent file should return error")
	}

	// Create an invalid PEM file
	invalidFile := "/tmp/invalid_cert.pem"
	defer os.Remove(invalidFile)

	err = os.WriteFile(invalidFile, []byte("invalid pem content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create invalid certificate file: %v", err)
	}

	_, err = LoadCertificate(invalidFile)
	if err == nil {
		t.Error("LoadCertificate() with invalid PEM should return error")
	}
}
