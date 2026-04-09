package probe

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"testing"
	"time"
)

func TestTLSConfig_Validate_Disabled(t *testing.T) {
	cfg := &TLSConfig{Enabled: false}
	if err := cfg.Validate(); err != nil {
		t.Fatalf("expected no error for disabled TLS, got %v", err)
	}
}

func TestTLSConfig_Validate_MismatchedKeyPair(t *testing.T) {
	cfg := &TLSConfig{Enabled: true, ClientCertFile: "cert.pem"}
	if err := cfg.Validate(); err == nil {
		t.Fatal("expected error when only cert is set without key")
	}
}

func TestTLSConfig_BuildCredentials_NoTLS(t *testing.T) {
	cfg := &TLSConfig{Enabled: false}
	creds, err := cfg.BuildCredentials()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if creds != nil {
		t.Fatal("expected nil credentials for disabled TLS")
	}
}

func TestTLSConfig_BuildCredentials_InsecureSkipVerify(t *testing.T) {
	cfg := &TLSConfig{Enabled: true, InsecureSkipVerify: true}
	creds, err := cfg.BuildCredentials()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if creds == nil {
		t.Fatal("expected non-nil credentials")
	}
}

func TestTLSConfig_BuildCredentials_WithCACert(t *testing.T) {
	caFile := writeSelfSignedCert(t)
	cfg := &TLSConfig{Enabled: true, CACertFile: caFile, InsecureSkipVerify: false}
	creds, err := cfg.BuildCredentials()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if creds == nil {
		t.Fatal("expected non-nil credentials")
	}
}

func TestTLSConfig_BuildCredentials_InvalidCAFile(t *testing.T) {
	cfg := &TLSConfig{Enabled: true, CACertFile: "/nonexistent/ca.pem"}
	_, err := cfg.BuildCredentials()
	if err == nil {
		t.Fatal("expected error for missing CA cert file")
	}
}

// writeSelfSignedCert generates a temporary self-signed PEM cert and returns its path.
func writeSelfSignedCert(t *testing.T) string {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("generating key: %v", err)
	}
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "test"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
	}
	certDER, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	if err != nil {
		t.Fatalf("creating cert: %v", err)
	}
	f, err := os.CreateTemp(t.TempDir(), "ca-*.pem")
	if err != nil {
		t.Fatalf("creating temp file: %v", err)
	}
	if err := pem.Encode(f, &pem.Block{Type: "CERTIFICATE", Bytes: certDER}); err != nil {
		t.Fatalf("encoding cert: %v", err)
	}
	f.Close()
	return f.Name()
}
