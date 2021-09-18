package paseto

import (
	"fmt"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

var PaseMaker = NewPasetoMaker("WrrTAEy8j5#&2&!8&iEF974Jh3#wNdz8")

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker(symmetricKey string) Maker {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker
}

// CreateToken creates a new token for a specific username and duration
func (maker *PasetoMaker) CreateToken(payloadGet PayloadGet) (string, error) {
	payload, err := NewPayload(payloadGet)
	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
