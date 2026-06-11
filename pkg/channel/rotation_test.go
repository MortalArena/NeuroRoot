package channel

import (
	"crypto/ed25519"
	"testing"
)

func TestKeyRotationOnRemove(t *testing.T) {
	ownerPub, ownerPriv, _ := ed25519.GenerateKey(nil)
	m1Pub, _, _ := ed25519.GenerateKey(nil)
	m2Pub, _, _ := ed25519.GenerateKey(nil)

	ownerDID := "did:ia:owner"
	m1DID := "did:ia:member1"
	m2DID := "did:ia:member2"

	cfg, oldKey, err := NewPrivateChannel("ch-1", ownerDID, ownerPriv, []string{m1DID, m2DID}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(oldKey) != 32 {
		t.Fatal("invalid key length")
	}

	memberPubs := map[string]ed25519.PublicKey{
		m1DID: m1Pub,
		m2DID: m2Pub,
	}

	newKey, err := cfg.RemoveMember(m1DID, ownerDID, ownerPriv, memberPubs)
	if err != nil {
		t.Fatal(err)
	}
	if string(newKey) == string(oldKey) {
		t.Fatal("key should rotate on member removal")
	}
	if cfg.KeyVersion != 2 {
		t.Fatalf("expected key version 2, got %d", cfg.KeyVersion)
	}
	if cfg.IsMember(m1DID) {
		t.Fatal("removed member should not be in list")
	}
	if _, ok := cfg.MemberKeys[m1DID]; ok {
		t.Fatal("removed member should not have encrypted key")
	}
	if err := cfg.VerifyConfigV2(ownerPub); err != nil {
		t.Fatalf("verify failed: %v", err)
	}
}
