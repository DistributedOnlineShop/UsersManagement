package token

import (
	"crypto/ed25519"
	"crypto/sha512"
	"fmt"

	"github.com/biscuit-auth/biscuit-go"
)

type KeyMaker struct {
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
}

func CreateKey(KeySeed string) (KeyMaker, error) {
	hash := sha512.Sum512([]byte(KeySeed))

	keyPair := KeyMaker{
		PrivateKey: ed25519.NewKeyFromSeed(hash[:32]),
	}

	keyPair.PublicKey = keyPair.PrivateKey.Public().(ed25519.PublicKey)

	return keyPair, nil
}

func (keypair *KeyMaker) CreateBiscuitToken(email, role string) ([]byte, *Payload, error) {
	builder := biscuit.NewBuilder(keypair.PrivateKey)

	payload, err := NewPayload(email, role)
	if err != nil {
		return nil, &Payload{}, fmt.Errorf("Failed to create payload: %w", err)
	}

	builder.AddAuthorityFact(biscuit.Fact{
		Predicate: biscuit.Predicate{
			Name: "user",
			IDs: []biscuit.Term{
				biscuit.String(payload.Email),
				biscuit.String(payload.Role),
			},
		},
	})
	builder.AddAuthorityFact(biscuit.Fact{
		Predicate: biscuit.Predicate{
			Name: "issued_at",
			IDs: []biscuit.Term{
				biscuit.String(payload.IssuedAt.String()),
			},
		},
	})
	builder.AddAuthorityFact(biscuit.Fact{
		Predicate: biscuit.Predicate{
			Name: "expired_at",
			IDs: []biscuit.Term{
				biscuit.String(payload.ExpiredAt.String()),
			},
		},
	})

	token, err := builder.Build()
	if err != nil {
		return nil, &Payload{}, fmt.Errorf("Failed to create token: %w", err)
	}

	serializedToken, err := token.Serialize()
	if err != nil {
		return nil, &Payload{}, fmt.Errorf("Failed to serialize token: %w", err)
	}

	return serializedToken, payload, nil
}

func (keypair *KeyMaker) VerifyBiscuitToken(serializedToken []byte, payload *Payload) (Payload, error) {
	token, err := biscuit.Unmarshal(serializedToken)
	if err != nil {
		return Payload{}, fmt.Errorf("Failed to parse token: %w", err)
	}

	verifier, err := token.Authorizer(keypair.PublicKey)
	if err != nil {
		return Payload{}, fmt.Errorf("Failed to create authorizer: %w", err)
	}

	verifier.AddCheck(biscuit.Check{
		Queries: []biscuit.Rule{
			{
				Head: biscuit.Predicate{
					Name: "allow",
					IDs:  []biscuit.Term{},
				},
				Body: []biscuit.Predicate{
					{
						Name: "user",
						IDs: []biscuit.Term{
							biscuit.String(payload.Email),
							biscuit.String(payload.Role),
						},
					},
					{
						Name: "issued_at",
						IDs: []biscuit.Term{
							biscuit.String(payload.IssuedAt.String()),
						},
					},
					{
						Name: "expired_at",
						IDs: []biscuit.Term{
							biscuit.String(payload.ExpiredAt.String()),
						},
					},
				},
			},
		},
	})

	verifier.AddPolicy(biscuit.Policy{
		Queries: []biscuit.Rule{
			{
				Head: biscuit.Predicate{
					Name: "allow",
					IDs:  []biscuit.Term{},
				},
				Body: []biscuit.Predicate{
					{
						Name: "user",
						IDs: []biscuit.Term{
							biscuit.String(payload.Email),
							biscuit.String(payload.Role),
						},
					},
					{
						Name: "issued_at",
						IDs: []biscuit.Term{
							biscuit.String(payload.IssuedAt.String()),
						},
					},
					{
						Name: "expired_at",
						IDs: []biscuit.Term{
							biscuit.String(payload.ExpiredAt.String()),
						},
					},
				},
			},
		},
	})

	err = verifier.Authorize()
	if err != nil {
		return Payload{}, fmt.Errorf("verification failed: %w", err)
	}

	return Payload{
		Email:     payload.Email,
		Role:      payload.Role,
		IssuedAt:  payload.IssuedAt,
		ExpiredAt: payload.ExpiredAt,
	}, nil
}
