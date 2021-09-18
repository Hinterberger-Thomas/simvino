package paseto

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type PayloadGet struct {
	Role   string `json:"role"`
	UserId uint32 `json:"email"`
}

// Payload contains the payload data of the token
type Payload struct {
	TokenID  uuid.UUID `json:"tokenID"`
	Role     string    `json:"role"`
	UserId   uint32    `json:"email"`
	IssuedAt time.Time `json:"issued_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(payloadGet PayloadGet) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		TokenID:  tokenID,
		UserId:   payloadGet.UserId,
		Role:     payloadGet.Role,
		IssuedAt: time.Now(),
	}
	return payload, nil
}
