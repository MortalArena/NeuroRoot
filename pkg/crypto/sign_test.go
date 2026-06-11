package crypto

import (
	"crypto/ed25519"
	"testing"
)

func TestSignAndVerify(t *testing.T) {
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatal(err)
	}
	payload := "test|payload|123"
	sig, err := SignPayloadHex(priv, DomainIdentity, payload)
	if err != nil {
		t.Fatal(err)
	}
	if err := VerifyPayloadHex(pub, DomainIdentity, payload, sig); err != nil {
		t.Fatalf("verify failed: %v", err)
	}
}

func TestDomainSeparation(t *testing.T) {
	pub, priv, _ := ed25519.GenerateKey(nil)
	payload := "same|payload"
	sig, _ := SignPayloadHex(priv, DomainIdentity, payload)
	// توقيع من domain مختلف يجب أن يفشل
	if err := VerifyPayloadHex(pub, DomainRevocation, payload, sig); err == nil {
		t.Fatal("cross-domain replay should fail")
	}
}

func TestRandomNonce(t *testing.T) {
	n1, err := RandomNonce()
	if err != nil {
		t.Fatal(err)
	}
	n2, err := RandomNonce()
	if err != nil {
		t.Fatal(err)
	}
	if len(n1) != 16 || len(n2) != 16 {
		t.Fatal("nonce must be 16 bytes")
	}
}
