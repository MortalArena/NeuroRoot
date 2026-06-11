package acp

import (
	"crypto/ed25519"
	"encoding/json"
	"testing"
)

func TestEnvelopeSignVerify(t *testing.T) {
	pub, priv, _ := ed25519.GenerateKey(nil)
	input, _ := json.Marshal(map[string]string{"text": "hello"})

	env, err := NewTaskRequest("did:ia:sender", "did:ia:receiver", TaskEcho, input, "", "req-1", priv)
	if err != nil {
		t.Fatal(err)
	}
	if err := VerifyEnvelope(env, pub); err != nil {
		t.Fatalf("verify failed: %v", err)
	}
}

func TestRouterEcho(t *testing.T) {
	r := NewRouter()
	ctx := t.Context()
	input, _ := json.Marshal(map[string]string{"text": "test"})
	out, err := r.Handle(ctx, TaskEcho, input)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]string
	json.Unmarshal(out, &result)
	if result["echo"] != "test" {
		t.Fatalf("expected echo test, got %s", result["echo"])
	}
}
