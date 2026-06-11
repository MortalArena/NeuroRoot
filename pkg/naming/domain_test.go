package naming

import (
	"crypto/ed25519"
	"testing"
	"time"
)

func TestNormalizeDomainName(t *testing.T) {
	tests := []struct {
		input   string
		wantErr bool
	}{
		{"example.ia", false},
		{"my-site.ia", false},
		{"example.com", true},
		{"", true},
		{"-bad.ia", true},
	}
	for _, tt := range tests {
		_, err := NormalizeDomainName(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("NormalizeDomainName(%q) err=%v wantErr=%v", tt.input, err, tt.wantErr)
		}
	}
}

func TestDomainRecordSignatures(t *testing.T) {
	_, founderPriv, _ := ed25519.GenerateKey(nil)
	ownerPub, ownerPriv, _ := ed25519.GenerateKey(nil)
	ownerDID := "did:ia:testowner123456"

	exp := time.Now().Add(365 * 24 * time.Hour).Unix()
	rec, err := NewDomainRecord("test.ia", ownerDID, "target", "did", exp, founderPriv)
	if err != nil {
		t.Fatal(err)
	}
	founderPub := founderPriv.Public().(ed25519.PublicKey)
	if err := rec.VerifyFounderSig(founderPub); err != nil {
		t.Fatalf("founder sig failed: %v", err)
	}
	if err := rec.SignOwner(ownerPriv); err != nil {
		t.Fatal(err)
	}
	if err := rec.VerifyOwnerSig(ownerPub); err != nil {
		t.Fatalf("owner sig failed: %v", err)
	}
}
