package sshutils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

type GenerateSSH struct {
	bitSize         int
	privateFileName string
	publicFileName  string

	privateFile []byte
	publicFile  []byte
}

func NewGenerateSSH(name, password string) *GenerateSSH {
	return &GenerateSSH{
		bitSize:         4096,
		privateFileName: fmt.Sprintf("id_rsa_%s", strings.ToLower(name)),
		publicFileName:  fmt.Sprintf("id_rsa_%s.pub", strings.ToLower(name)),
	}
}

func (u *GenerateSSH) GenerateSSHKey() error {
	// Generate Private KEY
	private, err := u.generatePrivateKey()
	if err != nil {
		return err
	}

	u.privateFile = u.encodePrivateKeyToPEM(private)

	// Generate Public Key
	public, err := u.generatePublicKey(&private.PublicKey)
	if err != nil {
		return err
	}

	u.publicFile = public

	return nil
}

func (u *GenerateSSH) GetPrivateKey() (privateName string, privateKey []byte) {
	return u.privateFileName, u.privateFile
}

func (u *GenerateSSH) GetPublicKey() (publicName string, publicKey []byte) {
	return u.publicFileName, u.publicFile
}

// generatePrivateKey creates an RSA Private Key of specified byte size
func (u *GenerateSSH) generatePrivateKey() (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, u.bitSize)
	if err != nil {
		return nil, err
	}

	if err := privateKey.Validate(); err != nil {
		return nil, err
	}

	return privateKey, nil
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func (u *GenerateSSH) encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privateDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privateBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateDER,
	}

	// Private key in PEM format
	return pem.EncodeToMemory(&privateBlock)
}

func (u *GenerateSSH) generatePublicKey(privateKey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privateKey)
	if err != nil {
		return nil, err
	}

	return ssh.MarshalAuthorizedKey(publicRsaKey), nil
}
