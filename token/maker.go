package token

type Maker interface {
	CreateBiscuitToken(email, role string) ([]byte, *Payload, error)
	VerifyBiscuitToken(serializedToken []byte, payload *Payload) (Payload, error)
}
